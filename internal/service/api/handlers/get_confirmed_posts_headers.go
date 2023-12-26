package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/types"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetConfirmedPostsHeaders(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPostsHeaders(r)
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	postsQ := ctx.PostsQ(r).New().Limit(request.Limit)
	if request.Sorting == types.DescSorting {
		postsQ.WhereIDLessThan(request.Cursor).OrderByID(types.DescSorting)
	} else {
		postsQ.WhereIDGreaterThan(request.Cursor).OrderByID(types.AscSorting)
	}
	if request.AuthorID != nil {
		postsQ.JoinUsersOnID().FilterByUsersAuthorID(*request.AuthorID)
	}
	postsHeaders, err := postsQ.FilterByStatus(types.ConfirmedPostStatus).SelectHeaders()
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
	response, err := responses.ComposePostHeaderListResponse(postsHeaders, r)
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
