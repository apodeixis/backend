package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
)

func DeletePostFromStarred(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeletePostFromStarred(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = ctx.StarredPostsQ(r).New().Transaction(func() error {
		err = ctx.StarredPostsQ(r).New().
			FilterByUserID(request.UserID).FilterByPostID(request.PostID).Delete()
		return err
	})
	if errors.Is(errors.Cause(err), postgres.ErrNoSuchStarredPost) {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
