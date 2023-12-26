package parameters

import (
	"net/http"
	"strconv"
)

func FetchQueryAuthorID(r *http.Request) (int64, error) {
	const queryKey = "author[id]"
	value := r.URL.Query().Get(queryKey)
	if value == "" {
		return -1, ErrQueryParamIsNotPresent
	}
	limit, err := strconv.Atoi(value)
	if limit < 0 {
		return -1, ErrUnsupportedValue
	}
	return int64(limit), err
}
