package parameters

import (
	"net/http"
	"strconv"
)

func FetchQueryCursor(r *http.Request) (int64, error) {
	const queryKey = "cursor"
	value := r.URL.Query().Get(queryKey)
	if value == "" {
		return -1, ErrQueryParamIsNotPresent
	}
	number, err := strconv.Atoi(value)
	if number < 0 {
		return -1, ErrUnsupportedValue
	}
	return int64(number), err
}
