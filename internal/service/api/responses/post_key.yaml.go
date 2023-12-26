package responses

import (
	"strconv"

	"github.com/apodeixis/backend/resources"
)

type PostKeyResponse struct {
	Data resources.PostKey `json:"data"`
}

func ComposePostKey(postID int64) *PostKeyResponse {
	return &PostKeyResponse{
		Data: resources.PostKey{
			Id:   strconv.FormatInt(postID, 10),
			Type: "post",
		},
	}
}
