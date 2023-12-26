package sender

import (
	"context"

	"github.com/apodeixis/backend/internal/data"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

func hashPost(post *data.Post) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	payload := []byte(post.Title + post.Body)
	hash.Write(payload)
	hashBytes := hash.Sum(nil)
	var hashArray [32]byte
	copy(hashArray[:], hashBytes)
	return hashArray
}

func (s *Service) composeTransactBindOpts(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := s.evmChainConfig.Client.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(s.evmChainConfig.SenderPrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	opts.GasPrice, err = s.evmChainConfig.Client.SuggestGasPrice(ctx)
	return opts, nil
}
