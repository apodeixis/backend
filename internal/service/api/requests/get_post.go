package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
)

type GetPostRequest struct {
	PostID int64
}

func NewGetPost(r *http.Request) (*GetPostRequest, error) {
	postID, err := parameters.FetchPathID(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch post id from path"),
		}
	}
	return &GetPostRequest{PostID: postID}, err
}
