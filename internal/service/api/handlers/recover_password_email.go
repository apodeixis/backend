package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"github.com/apodeixis/backend/internal/types"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RecoverPasswordEmail(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewEmail(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		errObjects := problems.BadRequest(err)
		if strings.Contains(err.Error(), is.ErrEmail.Error()) {
			if len(errObjects) == 1 {
				errObjects[0].Code = types.ErrInvalidEmail
			}
		}
		ape.RenderErr(w, errObjects...)
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByEmail(request.Data.Attributes.Email).Get()
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		log.Error("not found user")
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundUser))
		return
	}
	if user.OAuth2User {
		log.Warn("unable to recover password: used OAuth2")
		ape.RenderErr(w, problems.UnprocessableEntity(types.ErrUnprocessableEntityOAuth2User))
		return
	}
	token, err := helpers.GenerateToken(ctx.PasswordRecoveryConfig(r).Secret)
	if err != nil {
		log.WithError(err).Error("failed to generate token string")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	expiresAt := time.Now().Add(ctx.PasswordRecoveryConfig(r).TokenLife).UTC()
	q := ctx.PasswordRecoveryTokensQ(r).New()
	err = q.Transaction(func() error {
		_, err = q.Create(data.PasswordRecoveryToken{
			UserID:         user.ID,
			Token:          token,
			TokenExpiresAt: &expiresAt,
		})
		return err
	})
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	payload := helpers.ComposeRecoveryPayload(user.Name, token, expiresAt, r)
	err = ctx.Notificator(r).SendRecoveryEmail(user.Email, payload)
	if err != nil {
		log.WithError(err).Error("failed to send password recovery email")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeEmail(user.Email)
	ape.Render(w, response)
}
