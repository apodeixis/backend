package handlers

import (
	"net/http"
	"time"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"github.com/apodeixis/backend/internal/types"
)

func RecoverPassword(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewRecoverPassword(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	recovery, errObject := validateRecoveryToken(request.Data.Attributes.Token, r)
	if errObject != nil {
		ape.RenderErr(w, errObject)
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByID(recovery.UserID).Get()
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	password, err := helpers.HashPassword(request.Data.Attributes.NewPassword)
	if err != nil {
		log.WithError(err).Error("failed to hash password")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	user.Password = &password
	usersQ := ctx.UsersQ(r).New()
	err = usersQ.Transaction(func() error {
		user, err = usersQ.FilterByID(user.ID).Update(*user)
		if err != nil {
			return err
		}
		return ctx.PasswordRecoveryTokensQ(r).New().FilterByID(recovery.ID).Delete()
	})
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeUser(user)
	ape.Render(w, response)
}

func validateRecoveryToken(token string, r *http.Request) (*data.PasswordRecoveryToken, *jsonapi.ErrorObject) {
	log := ctx.Log(r)

	recovery, err := ctx.PasswordRecoveryTokensQ(r).New().FilterByToken(token).Get()
	if err != nil {
		log.Error(err)
		return nil, problems.InternalError()
	}
	if recovery == nil {
		log.Error("password recovery token not found")
		return nil, problems.NotFound(types.ErrNotFoundPasswordRecoveryToken)
	}
	// Needed to ensure that the password recovery token in db has not been forged:
	err = helpers.VerifyToken(recovery.Token, ctx.PasswordRecoveryConfig(r).Secret)
	if err != nil {
		log.WithError(err).Error("invalid password recovery token")
		return nil, problems.Unauthorized(types.ErrInvalidPasswordRecoveryToken)
	}
	if recovery.TokenExpiresAt.UTC().Before(time.Now().UTC()) {
		log.Error("password recovery token expired")
		return nil, problems.Unauthorized(types.ErrExpiredPasswordRecoveryToken)
	}
	return recovery, nil
}
