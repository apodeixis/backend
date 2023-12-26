package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

// TooManyRequests composes JSON:API error object with
// optional application-specific error code.
func TooManyRequests(code ...string) *jsonapi.ErrorObject {
	errorObject := &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusTooManyRequests),
		Status: fmt.Sprintf("%d", http.StatusTooManyRequests),
	}
	if len(code) == 1 {
		errorObject.Code = code[0]
	}
	if len(code) > 1 {
		panic(errors.New("code cannot be more than one value"))
	}
	return errorObject
}
