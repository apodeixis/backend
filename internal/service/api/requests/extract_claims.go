package requests

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers/jwt"
	"github.com/apodeixis/backend/internal/service/api/parameters"
)

func ExtractClaimsFromAuthHeader(r *http.Request) (*jwt.Claims, error) {
	token, err := parameters.FetchHeaderAuthorization(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch header authorization")
	}
	claims, err := jwt.ExtractClaims(token, ctx.JwtConfig(r).Secret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract claims")
	}
	return claims, nil
}
