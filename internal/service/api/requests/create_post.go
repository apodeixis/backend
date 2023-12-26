package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/resources"
)

type CreatePostRequest struct {
	Data   resources.CreatePost
	UserID int64
}

func NewCreatePost(r *http.Request) (*CreatePostRequest, error) {
	request := new(CreatePostRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	claims, err := ExtractClaimsFromAuthHeader(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract claims from auth header"),
		}
	}
	request.UserID = claims.OwnerId
	return request, request.validate()
}

func (r *CreatePostRequest) validate() error {
	return ozzo.Errors{
		"title": ozzo.Validate(r.Data.Attributes.Title, ozzo.Required, ozzo.Length(1, 512)),
		"body":  ozzo.Validate(r.Data.Attributes.Body, ozzo.Required, ozzo.Length(1, 2048)),
	}.Filter()
}
