package sender

import (
	"github.com/apodeixis/backend/internal/config"
	"github.com/apodeixis/backend/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type Service struct {
	log               *logan.Entry
	postsQ            data.Posts
	postTransactionsQ data.PostTransactions
	evmChainConfig    *config.EvmChainConfig
	newPosts          chan data.Post
	cursor            int64
}

type Opts struct {
	Log               *logan.Entry
	PostsQ            data.Posts
	PostTransactionsQ data.PostTransactions
	EvmChainConfig    *config.EvmChainConfig
}

func New(opts *Opts) *Service {
	return &Service{
		log:               opts.Log,
		postsQ:            opts.PostsQ,
		postTransactionsQ: opts.PostTransactionsQ,
		evmChainConfig:    opts.EvmChainConfig,
		newPosts:          make(chan data.Post, selectionLimit),
	}
}
