package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth2 "google.golang.org/api/oauth2/v2"
)

const oAuth2GoogleConfigKey = "oauth2_google"

var googleScopes = []string{
	goauth2.UserinfoEmailScope,
	goauth2.UserinfoProfileScope,
}

type oAuth2GoogleConfig struct {
	RedirectURL  string `figure:"redirect_url,required"`
	ClientID     string `figure:"client_id,required"`
	ClientSecret string `figure:"client_secret,required"`
}

func (c *config) OAuth2GoogleConfig() *oauth2.Config {
	return c.oAuth2GoogleConfig.Do(func() interface{} {
		config := new(oAuth2GoogleConfig)
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, oAuth2GoogleConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out oauth2_google config"))
		}

		return &oauth2.Config{
			RedirectURL:  config.RedirectURL,
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Scopes:       googleScopes,
			Endpoint:     google.Endpoint,
		}
	}).(*oauth2.Config)
}

type OAuth2GoogleStateConfig struct {
	StateSecret string
	StateLife   time.Duration
}

type oAuth2GoogleStateConfig struct {
	StateSecret string `figure:"state_secret,required"`
	StateLife   string `figure:"state_life"`
}

func (c *config) OAuth2GoogleStateConfig() *OAuth2GoogleStateConfig {
	return c.oAuth2GoogleStateConfig.Do(func() interface{} {
		const (
			defaultStateLife = "10m5s"
		)
		config := &oAuth2GoogleStateConfig{
			StateLife: defaultStateLife,
		}
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, oAuth2GoogleConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out oauth2_google state config"))
		}
		stateLife, err := time.ParseDuration(config.StateLife)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse oauth2_google state life duration"))
		}
		return &OAuth2GoogleStateConfig{
			StateSecret: config.StateSecret,
			StateLife:   stateLife,
		}
	}).(*OAuth2GoogleStateConfig)
}
