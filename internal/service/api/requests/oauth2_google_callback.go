package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/resources"
)

type OAuth2GoogleCallbackRequest struct {
	Data resources.OAuth2Callback
}

func NewOAuth2GoogleCallbackRequest(r *http.Request) (*OAuth2GoogleCallbackRequest, error) {
	request := new(OAuth2GoogleCallbackRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *OAuth2GoogleCallbackRequest) validate() error {
	return ozzo.Errors{
		"code":  ozzo.Validate(&r.Data.Attributes.Code, ozzo.Required),
		"state": ozzo.Validate(&r.Data.Attributes.State, ozzo.Required),
	}.Filter()
}
