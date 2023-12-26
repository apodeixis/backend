package requests

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/service/api/parameters"
	"github.com/apodeixis/backend/internal/types"
)

type cursorPagination struct {
	Limit   uint64
	Cursor  int64
	Sorting types.Sorting
}

func newCursorPagination(r *http.Request) (*cursorPagination, error) {
	const (
		defaultLimit   = 15
		defaultCursor  = 2147483647 // serial postgres max value
		defaultSorting = types.DescSorting
	)

	limit, err := parameters.FetchQueryLimit(r)
	if err != nil && !errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		return nil, errors.Wrap(err, "failed to fetch query limit")
	}
	if errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		limit = defaultLimit
	}

	cursor, err := parameters.FetchQueryCursor(r)
	if err != nil && !errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		return nil, errors.Wrap(err, "failed to fetch query cursor")
	}
	if errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		cursor = defaultCursor
	}

	sorting, err := parameters.FetchQuerySorting(r)
	if err != nil && !errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		return nil, errors.Wrap(err, "failed to fetch query sorting")
	}
	if errors.Is(err, parameters.ErrQueryParamIsNotPresent) {
		sorting = defaultSorting
	}
	return &cursorPagination{
		Limit:   limit,
		Cursor:  cursor,
		Sorting: sorting,
	}, nil
}
