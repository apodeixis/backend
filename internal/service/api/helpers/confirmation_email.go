package helpers

import (
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"

	"github.com/pkg/errors"
)

func SendConfirmationEmail(user *data.User, r *http.Request) error {
	token, err := GenerateToken(ctx.EmailVerificationConfig(r).Secret)
	if err != nil {
		return errors.Wrap(err, "failed to generate token string")
	}
	expiresAt := time.Now().Add(ctx.EmailVerificationConfig(r).TokenLife).UTC()
	q := ctx.EmailVerificationTokensQ(r).New()
	err = q.Transaction(func() error {
		_, err = q.Create(data.EmailVerificationToken{
			UserID:         user.ID,
			Token:          token,
			TokenExpiresAt: &expiresAt,
		})
		return err
	})
	if err != nil {
		return err
	}
	payload := ComposeConfirmationPayload(user.Name, token, expiresAt, r)
	err = ctx.Notificator(r).SendConfirmationEmail(user.Email, payload)
	return errors.Wrap(err, "failed to send confirmation email")
}
