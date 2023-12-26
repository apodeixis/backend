package responses

import (
	"strconv"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/resources"
)

type PostWithUserResponse struct {
	Data     resources.Post                               `json:"data"`
	Included resources.CreatePost201ResponseIncludedInner `json:"included"`
}

func ComposePostWithUser(post *data.Post, user *data.User, starred bool) *PostWithUserResponse {
	var txTimestamp int64
	if post.TxTimestamp != nil {
		txTimestamp = post.TxTimestamp.Unix()
	}

	return &PostWithUserResponse{
		Data: resources.Post{
			Id:   strconv.FormatInt(post.ID, 10),
			Type: "post",
			Attributes: resources.PostAllOfAttributes{
				Body:        post.Body,
				Title:       post.Title,
				TxHash:      post.TxHash,
				TxTimestamp: &txTimestamp,
				Status:      string(post.Status),
				Starred:     starred,
			},
			Relationships: &resources.PostAllOfRelationships{
				Author: *resources.NewPostAllOfRelationshipsAuthor(resources.UserKey{
					Id:   strconv.FormatInt(user.ID, 10),
					Type: "user",
				}),
			},
		},
		Included: resources.CreatePost201ResponseIncludedInner{
			User: convertToUserResource(user),
		},
	}
}
