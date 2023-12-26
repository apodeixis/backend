package requests

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
)

type GetConfirmedPostsHeadersRequest struct {
	cursorPagination

	AuthorID *int64
}

func NewGetPostsHeaders(r *http.Request) (*GetConfirmedPostsHeadersRequest, error) {
	cursorPagination, err := newCursorPagination(r)
	if err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to fetch pagination params"),
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

	return &GetConfirmedPostsHeadersRequest{
		cursorPagination: *cursorPagination,

		AuthorID: authorID,
	}, nil
}
