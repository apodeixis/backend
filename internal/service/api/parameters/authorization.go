package parameters

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrNoAuthHeader      = errors.New("no auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
)

func FetchHeaderAuthorization(r *http.Request) (string, error) {
	const (
		headerKey  = "Authorization"
		authScheme = "Bearer"
	)
	authHeader := r.Header.Get(headerKey)
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	if !strings.HasPrefix(authHeader, authScheme) {
		return "", ErrInvalidAuthHeader
	}
	token := strings.TrimPrefix(authHeader, authScheme)
	token = strings.TrimSpace(token)
	return token, nil
}
