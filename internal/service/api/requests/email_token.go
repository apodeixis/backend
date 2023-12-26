package requests

import (
	"encoding/json"
	"net/http"

	"github.com/apodeixis/backend/resources"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type EmailTokenRequest struct {
	Data resources.EmailToken
}

func NewEmailToken(r *http.Request) (*EmailTokenRequest, error) {
	request := new(EmailTokenRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *EmailTokenRequest) validate() error {
	return ozzo.Errors{
		"token": ozzo.Validate(r.Data.Attributes.Token, ozzo.Required),
	}.Filter()
}
