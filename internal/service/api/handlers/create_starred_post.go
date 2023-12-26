package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
)

func CreateStarredPost(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateStarredPost(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = ctx.StarredPostsQ(r).New().Transaction(func() error {
		_, err = ctx.StarredPostsQ(r).New().Create(data.StarredPost{
			UserID: request.UserID,
			PostID: request.PostID,
		})
		return err
	})
	switch err := errors.Cause(err); {
	case errors.Is(err, postgres.ErrNoSuchPost):
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.NotFound())
	case errors.Is(err, postgres.ErrNoSuchUser):
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.NotFound())
	case errors.Is(err, postgres.ErrStarredPostAlreadyExists):
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.Conflict())
	case err == nil:
		response := responses.ComposePostKey(request.PostID)
		w.WriteHeader(http.StatusCreated)
		ape.Render(w, response)
	default:
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
	}
}
