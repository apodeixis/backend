package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPost(r)
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	post, err := ctx.PostsQ(r).New().FilterByID(request.PostID).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if post == nil {
		ctx.Log(r).Error("post not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByID(post.UserID).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	starred := false
	claims, _ := requests.ExtractClaimsFromAuthHeader(r)
	if claims != nil {
		starredForUserID := claims.OwnerId
		starredPost, err := ctx.StarredPostsQ(r).New().
			FilterByPostID(post.ID).FilterByUserID(starredForUserID).Get()
		if err != nil {
			ctx.Log(r).Error(err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		starred = starredPost != nil
	}
	w.WriteHeader(http.StatusOK)
	response := responses.ComposePostWithUser(post, user, starred)
	ape.Render(w, response)
}
