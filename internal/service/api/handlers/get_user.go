package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUser(r)
	if err != nil {
		ctx.Log(r).Error(err)
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
		ctx.Log(r).Error("user not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	response := responses.ComposeUser(user)
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
