package parameters

import (
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrUnsupportedValue       = errors.New("unsupported value")
	ErrQueryParamIsNotPresent = errors.New("query param is not present")
)

func FetchQueryLimit(r *http.Request) (uint64, error) {
	const queryKey = "limit"
	value := r.URL.Query().Get(queryKey)
	if value == "" {
		return 0, ErrQueryParamIsNotPresent
	}
	limit, err := strconv.Atoi(value)
	if limit < 1 {
		return 0, ErrUnsupportedValue
	}
	return uint64(limit), err
}
