package parameters

import (
	"fmt"
	"net/http"

	"github.com/apodeixis/backend/internal/types"
)

func FetchQueryPostStatus(r *http.Request) (types.PostStatus, error) {
	const queryKey = "post[status]"
	switch r.URL.Query().Get(queryKey) {
	case "":
		return "", ErrQueryParamIsNotPresent
	case fmt.Sprint(types.NewPostStatus):
		return types.NewPostStatus, nil
	case fmt.Sprint(types.PendingPostStatus):
		return types.PendingPostStatus, nil
	case fmt.Sprint(types.ConfirmedPostStatus):
		return types.ConfirmedPostStatus, nil
	case fmt.Sprint(types.FailedPostStatus):
		return types.FailedPostStatus, nil
	}
	return "", ErrUnsupportedValue
}
