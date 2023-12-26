package requests

import (
	"encoding/json"
	"net/http"

	"github.com/apodeixis/backend/resources"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
)

type SignUpRequest struct {
	Data resources.SignUpUser
}

func NewSignUp(r *http.Request) (*SignUpRequest, error) {
	request := new(SignUpRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *SignUpRequest) validate() error {
	return ozzo.Errors{
		"email":    ozzo.Validate(r.Data.Attributes.Email, ozzo.Required, is.Email),
		"password": ozzo.Validate(r.Data.Attributes.Password, ozzo.Required, ozzo.Length(8, 32)),
		"name":     ozzo.Validate(r.Data.Attributes.Name, ozzo.Required, ozzo.Length(1, 255)),
	}.Filter()
}
