package config

import (
	"net/url"

	"github.com/apodeixis/backend/internal/connectors/notificator"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const notificatorConfigKey = "notificator"

type notificatorConfig struct {
	Endpoint *url.URL `fig:"endpoint,required"`
}

func (c *config) Notificator() notificator.Connector {
	return c.notificator.Do(func() interface{} {
		var disabledConfig struct {
			Disabled bool `fig:"disabled"`
		}
		err := figure.
			Out(&disabledConfig).
			From(kv.MustGetStringMap(c.getter, notificatorConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator disabled configuration"))
		}
		if disabledConfig.Disabled {
			return notificator.NewDisabledConnector(c.Log())
		}
		config := new(notificatorConfig)
		err = figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, notificatorConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator config"))
		}
		return notificator.NewConnector(config.Endpoint, c.Log())
	}).(notificator.Connector)
}
