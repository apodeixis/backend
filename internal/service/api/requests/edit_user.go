package requests

import (
	"encoding/json"
	"net/http"

	"github.com/apodeixis/backend/resources"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type EditUserRequest struct {
	Data   resources.EditUser
	UserID int64
}

func NewEditUser(r *http.Request) (*EditUserRequest, error) {
	request := new(EditUserRequest)
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

func (r *EditUserRequest) validate() error {
	return ozzo.Errors{
		"new_name": ozzo.Validate(r.Data.Attributes.NewName, ozzo.Length(1, 255)),
	}.Filter()
}
