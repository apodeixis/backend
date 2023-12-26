package handlers

import (
	"net/http"
	"strings"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/apodeixis/backend/internal/types"

	"github.com/google/jsonapi"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/requests"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSignUp(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		errObjects := problems.BadRequest(err)
		if strings.Contains(err.Error(), is.ErrEmail.Error()) {
			if len(errObjects) == 1 {
				errObjects[0].Code = types.ErrInvalidEmail
			}
		}
		ape.RenderErr(w, errObjects...)
		return
	}
	password, err := helpers.HashPassword(request.Data.Attributes.Password)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to hash password")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	user := &data.User{
		AuthorID: helpers.GenerateID(),
		Email:    request.Data.Attributes.Email,
		Name:     request.Data.Attributes.Name,
		Password: &password,
	}
	var errObject *jsonapi.ErrorObject
	err = ctx.UsersQ(r).New().Transaction(func() error {
		user, err = ctx.UsersQ(r).New().CreateUserFromSignUp(*user)
		if errors.Is(err, postgres.ErrUserWithSuchEmailAlreadyExists) {
			errObject = problems.Conflict(types.ErrConflictUserEmail)
			return err
		}
		return err
	})
	if err != nil {
		ctx.Log(r).Error(err)
		if errObject == nil {
			errObject = problems.InternalError()
		}
		ape.RenderErr(w, errObject)
		return
	}
	if err := helpers.SendConfirmationEmail(user, r); err != nil {
		// RenderErr ignore is intended
		ctx.Log(r).Error(err)
	}
	response := responses.ComposeUser(user)
	w.WriteHeader(http.StatusCreated)
	ape.Render(w, response)
}
