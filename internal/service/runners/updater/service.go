package updater

import (
	"context"
	"time"

	ctypes "github.com/apodeixis/backend/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/running"
)

const selectionLimit = 10

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Updater started")
	go running.WithBackOff(ctx, s.log, "selector",
		s.selectPendingPosts, 20*time.Second, 30*time.Second, time.Minute)

	go running.WithBackOff(ctx, s.log, "updater",
		s.updatePostsStatuses, 20*time.Second, 30*time.Second, time.Minute)
}

func (s *Service) selectPendingPosts(_ context.Context) error {
	selectInterval := time.Second * 3
	for {
		posts, err := s.postsQ.New().
			FilterByStatus(ctypes.PendingPostStatus).WhereIDGreaterThan(s.cursor).Limit(selectionLimit).Select()
		if err != nil {
			return err
		}
		for _, post := range posts {
			s.pendingPosts <- post
		}
		if len(posts) != 0 {
			cursor := posts[len(posts)-1].ID
			s.cursor = cursor
		}
		time.Sleep(selectInterval)
	}
}

func (s *Service) updatePostsStatuses(ctx context.Context) error {
	for post := range s.pendingPosts {
		postTx, err := s.postTransactionsQ.New().FilterByPostID(post.ID).Get()
		if err != nil {
			return err
		}
		receipt, err := bind.WaitMined(ctx, s.evmChainConfig.Client, (*types.Transaction)(postTx.Tx))
		if err != nil {
			return errors.Wrap(err, "failed to wait mined tx")
		}
		block, err := s.evmChainConfig.Client.BlockByNumber(ctx, receipt.BlockNumber)
		if err != nil {
			return errors.Wrap(err, "failed to get block by number")
		}
		timestamp := time.Unix(int64(block.Time()), 0)
		s.log.Infof("got receipt for post # %d, tx_hash %s", post.ID, *post.TxHash)
		if receipt.Status != types.ReceiptStatusSuccessful {
			post.Status = ctypes.FailedPostStatus
		} else {
			post.Status = ctypes.ConfirmedPostStatus
		}
		post.TxTimestamp = &timestamp
		err = s.postsQ.New().Transaction(func() error {
			_, err := s.postsQ.New().FilterByID(post.ID).Update(post)
			return err
		})
		if err != nil {
			return err
		}
		s.log.Infof("updated post # %d status to %s",
			post.ID, post.Status,
		)
	}
	return nil
}
