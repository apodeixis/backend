package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
	"github.com/apodeixis/backend/internal/types"
)

type GetUserPostsHeadersRequest struct {
	cursorPagination
	UserID int64

	Status *types.PostStatus
}

func NewGetUserPostsHeaders(r *http.Request) (*GetUserPostsHeadersRequest, error) {
	cursorPagination, err := newCursorPagination(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch pagination params"),
		}
	}

	claims, err := ExtractClaimsFromAuthHeader(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to extract claims from auth header"),
		}
	}

	statusValue, err := parameters.FetchQueryPostStatus(r)
	if err != nil && !errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch query post[status]"),
		}
	}
	status := &statusValue
	if errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		status = nil
	}

	return &GetUserPostsHeadersRequest{
		cursorPagination: *cursorPagination,
		UserID:           claims.OwnerId,

		Status: status,
	}, nil
}
