package config

import (
	"net/url"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const webConfigKey = "web"

type WebConfig struct {
	ConfirmationEndpoint *url.URL `fig:"confirmation_endpoint,required"`
	RecoveryEndpoint     *url.URL `fig:"recovery_endpoint,required"`
}

func (c *config) WebConfig() *WebConfig {
	return c.webConfig.Do(func() interface{} {
		config := new(WebConfig)
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, webConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out web config"))
		}
		return config
	}).(*WebConfig)
}
