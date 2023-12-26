package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const refreshCookieConfigKey = "refresh_cookie"

type RefreshCookieConfig struct {
	Name     string `figure:"name"`
	Path     string `figure:"path"`
	HttpOnly bool   `figure:"http_only"`
	Secure   bool   `figure:"secure"`
}

func (c *config) RefreshCookieConfig() *RefreshCookieConfig {
	return c.refreshCookieConfig.Do(func() interface{} {
		config := new(RefreshCookieConfig)
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, refreshCookieConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out refresh cookie config"))
		}
		return config
	}).(*RefreshCookieConfig)
}
