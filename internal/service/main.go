package service

import (
	"context"

	"github.com/apodeixis/backend/internal/config"
	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api"
	"github.com/apodeixis/backend/internal/service/runners/sender"
	"github.com/apodeixis/backend/internal/service/runners/updater"
	"github.com/apodeixis/backend/internal/types"
)

type Service struct {
	api           api.API
	mumbaiSender  *sender.Service
	mumbaiUpdater *updater.Service
}

func New(cfg config.Config) *Service {
	return &Service{
		api: api.NewAPI(cfg),
		mumbaiSender: sender.New(&sender.Opts{
			Log:               cfg.Log(),
			PostsQ:            postgres.NewPostsQ(cfg.DB()),
			PostTransactionsQ: postgres.NewPostTransactionsQ(cfg.DB()),
			EvmChainConfig:    cfg.EvmChainConfig(types.MumbaiEVMChain),
		}),
		mumbaiUpdater: updater.New(&updater.Opts{
			Log:               cfg.Log(),
			PostsQ:            postgres.NewPostsQ(cfg.DB()),
			PostTransactionsQ: postgres.NewPostTransactionsQ(cfg.DB()),
			EvmChainConfig:    cfg.EvmChainConfig(types.MumbaiEVMChain),
		}),
	}
}

func (s *Service) Run(ctx context.Context) error {
	s.mumbaiSender.Run(ctx)
	s.mumbaiUpdater.Run(ctx)
	return s.api.Start()
}
