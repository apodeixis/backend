package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/types"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetUserStarredPostsHeaders(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetStarredPostsHeaders(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	postsQ := ctx.PostsQ(r).New()
	if request.Sorting == types.DescSorting {
		postsQ.WhereIDLessThan(request.Cursor).OrderByID(types.DescSorting)
	} else {
		postsQ.WhereIDGreaterThan(request.Cursor).OrderByID(types.AscSorting)
	}
	if request.AuthorID != nil {
		postsQ.JoinUsersOnID().FilterByUsersAuthorID(*request.AuthorID)
	}
	postsHeaders, err := postsQ.Limit(request.Limit).
		JoinStarredPostsOnPostID().FilterByStarredPostsUserID(request.UserID).SelectHeaders()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(postsHeaders) == 0 {
		ctx.Log(r).Error("not found posts headers")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	response, err := responses.ComposeStarredPostHeaderListResponse(postsHeaders, ctx.UsersQ(r))
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
