package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
)

type GetUserRequest struct {
	UserID int64
}

func NewGetUser(r *http.Request) (*GetUserRequest, error) {
	userID, err := parameters.FetchPathID(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch user id from path"),
		}
	}
	return &GetUserRequest{UserID: userID}, err
}
