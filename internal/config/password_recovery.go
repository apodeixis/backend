package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const passwordRecoveryConfigKey = "password_recovery"

type PasswordRecoveryConfig struct {
	Secret    string
	TokenLife time.Duration
}

type passwordRecoveryConfig struct {
	Secret    string `figure:"secret,required"`
	TokenLife string `figure:"token_life"`
}

func (c *config) PasswordRecoveryConfig() *PasswordRecoveryConfig {
	return c.passwordRecoveryConfig.Do(func() interface{} {
		const (
			defaultTokenLife = "10m5s"
		)
		config := &passwordRecoveryConfig{
			TokenLife: defaultTokenLife,
		}
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, passwordRecoveryConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out password recovery config"))
		}
		tokenLife, err := time.ParseDuration(config.TokenLife)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse token life duration"))
		}
		return &PasswordRecoveryConfig{
			Secret:    config.Secret,
			TokenLife: tokenLife,
		}
	}).(*PasswordRecoveryConfig)
}
