package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
)

type GetStarredPostsHeadersRequest struct {
	cursorPagination
	UserID int64

	AuthorID *int64
}

func NewGetStarredPostsHeaders(r *http.Request) (*GetStarredPostsHeadersRequest, error) {
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

	authorIDValue, err := parameters.FetchQueryAuthorID(r)
	if err != nil && !errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch query user[id]"),
		}
	}
	authorID := &authorIDValue
	if errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		authorID = nil
	}

	return &GetStarredPostsHeadersRequest{
		cursorPagination: *cursorPagination,
		UserID:           claims.OwnerId,

		AuthorID: authorID,
	}, nil
}
