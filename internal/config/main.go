package config

import (
	"github.com/apodeixis/backend/internal/connectors/notificator"
	"github.com/apodeixis/backend/internal/types"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"golang.org/x/oauth2"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer

	JwtConfig() *JwtConfig
	OAuth2GoogleConfig() *oauth2.Config
	OAuth2GoogleStateConfig() *OAuth2GoogleStateConfig
	EvmChainConfig(chain types.EVMChain) *EvmChainConfig
	RefreshCookieConfig() *RefreshCookieConfig
	PasswordRecoveryConfig() *PasswordRecoveryConfig
	EmailVerificationConfig() *EmailVerificationConfig
	WebConfig() *WebConfig

	Notificator() notificator.Connector
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	getter kv.Getter

	oAuth2GoogleConfig      comfig.Once
	oAuth2GoogleStateConfig comfig.Once
	jwtConfig               comfig.Once
	refreshCookieConfig     comfig.Once
	passwordRecoveryConfig  comfig.Once
	emailVerificationConfig comfig.Once
	webConfig               comfig.Once

	notificator comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		Databaser:  pgdb.NewDatabaser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		getter:     getter,
	}
}
