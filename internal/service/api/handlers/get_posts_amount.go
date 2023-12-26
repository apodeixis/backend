package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/service/api/ctx"
)

func GetPostsAmount(w http.ResponseWriter, r *http.Request) {
	postsAmount, err := ctx.PostsQ(r).New().Count()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposePostsAmount(postsAmount)
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
