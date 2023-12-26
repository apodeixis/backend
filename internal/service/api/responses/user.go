package responses

import (
	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/resources"
)

type UserResponse struct {
	Data resources.User `json:"data"`
}

func ComposeUser(user *data.User) *UserResponse {
	return &UserResponse{
		Data: *convertToUserResource(user),
	}
}
