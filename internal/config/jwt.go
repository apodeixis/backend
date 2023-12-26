package config

import (
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const jwtConfigConfigKey = "jwt"

type JwtConfig struct {
	Secret      string
	RefreshLife time.Duration
	AccessLife  time.Duration
}

type jwtConfig struct {
	Secret      string `figure:"secret,required"`
	RefreshLife string `figure:"refresh_life"`
	AccessLife  string `figure:"access_life"`
}

func (c *config) JwtConfig() *JwtConfig {
	return c.jwtConfig.Do(func() interface{} {
		const (
			defaultRefreshLife = "6h"
			defaultAccessLife  = "10m5s"
		)
		config := &jwtConfig{
			RefreshLife: defaultRefreshLife,
			AccessLife:  defaultAccessLife,
		}
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, jwtConfigConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out jwt config"))
		}
		accessLife, err := time.ParseDuration(config.AccessLife)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse access life to duration"))
		}
		refreshLife, err := time.ParseDuration(config.RefreshLife)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse refresh life to duration"))
		}
		return &JwtConfig{
			Secret:      config.Secret,
			RefreshLife: refreshLife,
			AccessLife:  accessLife,
		}
	}).(*JwtConfig)
}
