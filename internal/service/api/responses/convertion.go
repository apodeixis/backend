package responses

import (
	"strconv"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/resources"
)

func ConvertToAuthTokenResource(access string, expiresAt int64) *resources.AuthToken {
	return resources.NewAuthToken("auth_token", resources.AuthTokenAllOfAttributes{
		Access:    access,
		ExpiresAt: expiresAt,
	})
}

func convertToUserResource(user *data.User) *resources.User {
	return &resources.User{
		Id:   strconv.FormatInt(user.ID, 10),
		Type: "user",
		Attributes: resources.UserAllOfAttributes{
			AuthorId:  int32(user.AuthorID),
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		},
	}
}

func convertPostHeaderToResource(postHeader *data.PostHeader, starred bool) *resources.PostHeader {
	var txTimestamp int64
	if postHeader.TxTimestamp != nil {
		txTimestamp = postHeader.TxTimestamp.Unix()
	}
	return &resources.PostHeader{
		Id:   strconv.FormatInt(postHeader.ID, 10),
		Type: "post_header",
		Attributes: resources.PostHeaderAllOfAttributes{
			Title:       postHeader.Title,
			TxHash:      postHeader.TxHash,
			TxTimestamp: &txTimestamp,
			Status:      string(postHeader.Status),
			Starred:     starred,
		},
	}
}
