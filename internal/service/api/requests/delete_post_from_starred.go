package requests

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/parameters"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type DeletePostFromStarredRequest struct {
	PostID int64
	UserID int64
}

func NewDeletePostFromStarred(r *http.Request) (*DeletePostFromStarredRequest, error) {
	postID, err := parameters.FetchPathID(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch post id from path"),
		}
	}
	claims, err := ExtractClaimsFromAuthHeader(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract claims from auth header"),
		}
	}
	return &DeletePostFromStarredRequest{
		PostID: postID,
		UserID: claims.OwnerId,
	}, nil
}
