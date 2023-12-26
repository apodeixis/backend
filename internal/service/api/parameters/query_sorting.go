package parameters

import (
	"fmt"
	"net/http"

	"github.com/apodeixis/backend/internal/types"
)

func FetchQuerySorting(r *http.Request) (types.Sorting, error) {
	const queryKey = "sorting"
	switch r.URL.Query().Get(queryKey) {
	case "":
		return "", ErrQueryParamIsNotPresent
	case fmt.Sprint(types.AscSorting):
		return types.AscSorting, nil
	case fmt.Sprint(types.DescSorting):
		return types.DescSorting, nil
	default:
		return "", ErrUnsupportedValue
	}
}
