package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/service/api/helpers"

	"github.com/apodeixis/backend/internal/service/api/ctx"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/service/api/requests"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewEditUser(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByID(request.UserID).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		ctx.Log(r).Error("not found user")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if request.Data.Attributes.NewName != nil {
		user.Name = *request.Data.Attributes.NewName
	}
	err = ctx.UsersQ(r).New().Transaction(func() error {
		user, err = ctx.UsersQ(r).New().FilterByID(user.ID).Update(*user)
		return err
	})
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	authToken, err := helpers.CreateAuthToken(user, r, w)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to create auth token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeUserWithAuthToken(user, authToken)
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
