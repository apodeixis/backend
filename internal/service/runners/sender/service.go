package sender

import (
	"context"
	"time"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/running"
)

const selectionLimit = 10

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Sender started")
	go running.WithBackOff(ctx, s.log, "selector",
		s.selectNewPosts, 20*time.Second, 30*time.Second, time.Minute)

	go running.WithBackOff(ctx, s.log, "transactor",
		s.transactNewPosts, 20*time.Second, 30*time.Second, time.Minute)
}

func (s *Service) selectNewPosts(_ context.Context) error {
	selectInterval := time.Second * 3
	for {
		posts, err := s.postsQ.New().
			FilterByStatus(types.NewPostStatus).WhereIDGreaterThan(s.cursor).Limit(selectionLimit).Select()
		if err != nil {
			return err
		}
		for _, post := range posts {
			s.newPosts <- post
		}
		if len(posts) != 0 {
			cursor := posts[len(posts)-1].ID
			s.cursor = cursor
		}
		time.Sleep(selectInterval)
	}
}

func (s *Service) transactNewPosts(ctx context.Context) error {
	for post := range s.newPosts {
		opts, err := s.composeTransactBindOpts(ctx)
		if err != nil {
			return err
		}
		postHash := hashPost(&post)
		tx, err := s.evmChainConfig.Contract.CreatePost(opts, post.UserID, postHash)
		if err != nil {
			return errors.Wrap(err, "failed to call contract create post method")
		}
		s.log.Infof("sent create post tx for post # %d", post.ID)
		err = s.postTransactionsQ.Transaction(func() error {
			post.Status = types.PendingPostStatus
			txHash := tx.Hash().String()
			post.TxHash = &txHash
			_, err := s.postsQ.New().FilterByID(post.ID).Update(post)
			if err != nil {
				return err
			}
			_, err = s.postTransactionsQ.New().Create(data.PostTransaction{
				PostID: post.ID,
				Tx:     (*types.RLPTransaction)(tx),
			})
			return err
		})
		if err != nil {
			return err
		}
		s.log.Infof("updated post # %d status to %s and tx_hash to %s",
			post.ID, post.Status, *post.TxHash,
		)
	}
	return nil
}
