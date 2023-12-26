package requests

import (
	"net/http"
	"regexp"

	"github.com/apodeixis/backend/internal/service/api/ctx"

	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/helpers/jwt"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type RefreshTokenRequest struct {
	Token string
}

func NewRefreshTokenRequest(r *http.Request) (*RefreshTokenRequest, error) {
	cookie, err := r.Cookie(ctx.RefreshCookieConfig(r).Name)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract refresh token cookie"),
		}
	}
	request := new(RefreshTokenRequest)
	request.Token = cookie.Value
	return request, request.validate()
}

func (r *RefreshTokenRequest) validate() error {
	tokenRegexp := regexp.MustCompile(jwt.TokenPattern)
	return ozzo.Errors{
		"token": ozzo.Validate(r.Token, ozzo.Required, ozzo.Match(tokenRegexp)),
	}.Filter()
}
