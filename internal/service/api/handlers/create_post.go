package handlers

import (
	"net/http"

	"github.com/google/jsonapi"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreatePost(r)
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	var (
		responseErr *jsonapi.ErrorObject
		response    interface{}
	)
	err = ctx.PostsQ(r).New().Transaction(func() error {
		post, err := ctx.PostsQ(r).New().Create(data.Post{
			UserID: request.UserID,
			Title:  request.Data.Attributes.Title,
			Body:   request.Data.Attributes.Body,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			responseErr = problems.InternalError()
			return err
		}
		user, err := ctx.UsersQ(r).New().FilterByID(post.UserID).Get()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			responseErr = problems.InternalError()
			return err
		}
		w.WriteHeader(http.StatusCreated)
		const starred = false

		response = responses.ComposePostWithUser(post, user, starred)
		return nil
	})
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, responseErr)
		return
	}
	ape.Render(w, response)
}
