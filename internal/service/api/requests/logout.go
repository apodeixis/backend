package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type LogoutRequest struct {
	UserID int64
}

func NewLogout(r *http.Request) (*LogoutRequest, error) {
	claims, err := ExtractClaimsFromAuthHeader(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract claims from auth header"),
		}
	}
	return &LogoutRequest{UserID: claims.OwnerId}, nil
}
