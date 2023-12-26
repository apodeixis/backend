package requests

import (
	"encoding/json"
	"net/http"

	"github.com/apodeixis/backend/resources"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
)

type EmailRequest struct {
	Data resources.Email
}

func NewEmail(r *http.Request) (*EmailRequest, error) {
	request := new(EmailRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *EmailRequest) validate() error {
	return ozzo.Errors{
		"email": ozzo.Validate(r.Data.Attributes.Email, ozzo.Required, is.Email),
	}.Filter()
}
