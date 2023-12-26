package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/resources"
)

type LoginRequest struct {
	Data resources.LoginUser
}

func NewLogin(r *http.Request) (*LoginRequest, error) {
	request := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *LoginRequest) validate() error {
	return ozzo.Errors{
		"email":    ozzo.Validate(r.Data.Attributes.Email, is.Email),
		"password": ozzo.Validate(r.Data.Attributes.Password, ozzo.Required, ozzo.Length(8, 32)),
	}.Filter()
}
