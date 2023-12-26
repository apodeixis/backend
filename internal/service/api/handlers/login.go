package handlers

import (
	"net/http"
	"strings"

	"github.com/apodeixis/backend/internal/types"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/service/api/helpers"

	"github.com/apodeixis/backend/internal/service/api/ctx"

	"github.com/apodeixis/backend/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewLogin(r)
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
	user, err := ctx.UsersQ(r).New().FilterByEmail(request.Data.Attributes.Email).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		ctx.Log(r).Error("not found user")
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundUser))
		return
	}
	if user.OAuth2User {
		ctx.Log(r).Error("oauth2 user")
		ape.RenderErr(w, problems.UnprocessableEntity(types.ErrUnprocessableEntityOAuth2User))
		return
	}
	err = helpers.CompareHashAndPassword(*user.Password, request.Data.Attributes.Password)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid password")
		ape.RenderErr(w, problems.Unauthorized(types.ErrInvalidPassword))
		return
	}
	var response interface{}
	if user.EmailVerified {
		authToken, err := helpers.CreateAuthToken(user, r, w)
		if err != nil {
			ctx.Log(r).WithError(err).Error("failed to create auth token")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		response = responses.ComposeUserWithAuthToken(user, authToken)
	} else {
		response = responses.ComposeUser(user)
	}
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
