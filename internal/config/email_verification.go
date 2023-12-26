package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

const emailVerificationConfigKey = "email_verification"

type EmailVerificationConfig struct {
	Secret    string
	TokenLife time.Duration
}

type emailVerificationConfig struct {
	Secret    string `figure:"secret,required"`
	TokenLife string `figure:"token_life"`
}

func (c *config) EmailVerificationConfig() *EmailVerificationConfig {
	return c.emailVerificationConfig.Do(func() interface{} {
		const (
			defaultTokenLife = "10m5s"
		)
		config := &emailVerificationConfig{
			TokenLife: defaultTokenLife,
		}
		err := figure.
			Out(config).
			From(kv.MustGetStringMap(c.getter, emailVerificationConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out email verification config"))
		}
		tokenLife, err := time.ParseDuration(config.TokenLife)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse token life duration"))
		}
		return &EmailVerificationConfig{
			Secret:    config.Secret,
			TokenLife: tokenLife,
		}
	}).(*EmailVerificationConfig)
}
