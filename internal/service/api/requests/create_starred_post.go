package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/resources"
)

type CreateStarredPostRequest struct {
	PostID int64
	UserID int64
}

func NewCreateStarredPost(r *http.Request) (*CreateStarredPostRequest, error) {
	request := new(resources.CreateStarredPostRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	postID, err := strconv.Atoi(request.Data.Id)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to convert post id to int"),
		}
	}
	claims, err := ExtractClaimsFromAuthHeader(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract claims from auth header"),
		}
	}
	return &CreateStarredPostRequest{
		PostID: int64(postID),
		UserID: claims.OwnerId,
	}, nil
}
