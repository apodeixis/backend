package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/resources"
)

type RecoverPasswordRequest struct {
	Data resources.RecoverPassword
}

func NewRecoverPassword(r *http.Request) (*RecoverPasswordRequest, error) {
	request := new(RecoverPasswordRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *RecoverPasswordRequest) validate() error {
	return ozzo.Errors{
		"token":        ozzo.Validate(r.Data.Attributes.Token, ozzo.Required),
		"new_password": ozzo.Validate(r.Data.Attributes.NewPassword, ozzo.Required, ozzo.Length(8, 32)),
	}.Filter()
}
