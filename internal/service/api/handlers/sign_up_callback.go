package handlers

import (
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"github.com/apodeixis/backend/internal/types"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func SignUpCallback(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewEmailToken(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	verification, errObject := validateVerificationToken(request.Data.Attributes.Token, r)
	if errObject != nil {
		ape.RenderErr(w, errObject)
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByID(verification.UserID).Get()
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user.EmailVerified {
		log.Error("email is already verified")
		ape.RenderErr(w, problems.Conflict(types.ErrEmailAlreadyVerified))
		return
	}
	user.EmailVerified = true
	usersQ := ctx.UsersQ(r).New()
	err = usersQ.Transaction(func() error {
		user, err = usersQ.FilterByID(user.ID).Update(*user)
		if err != nil {
			return err
		}
		return ctx.EmailVerificationTokensQ(r).New().FilterByID(verification.ID).Delete()
	})
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeUser(user)
	ape.Render(w, response)
}

func validateVerificationToken(token string, r *http.Request) (*data.EmailVerificationToken, *jsonapi.ErrorObject) {
	log := ctx.Log(r)

	verification, err := ctx.EmailVerificationTokensQ(r).New().FilterByToken(token).Get()
	if err != nil {
		log.Error(err)
		return nil, problems.InternalError()
	}
	if verification == nil {
		log.Error("email verification token not found")
		return nil, problems.NotFound(types.ErrNotFoundEmailVerificationToken)
	}
	// Needed to ensure that the email verification token in db has not been forged:
	err = helpers.VerifyToken(verification.Token, ctx.PasswordRecoveryConfig(r).Secret)
	if err != nil {
		log.WithError(err).Error("invalid email verification token")
		return nil, problems.Unauthorized(types.ErrInvalidEmailVerificationToken)
	}
	if verification.TokenExpiresAt.UTC().Before(time.Now().UTC()) {
		log.Error("email verification token expired")
		return nil, problems.Unauthorized(types.ErrExpiredEmailVerificationToken)
	}
	return verification, nil
}
