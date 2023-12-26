package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

// Unauthorized composes JSON:API error object with
// optional application-specific error code.
func Unauthorized(code ...string) *jsonapi.ErrorObject {
	errorObject := &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusUnauthorized),
		Status: fmt.Sprintf("%d", http.StatusUnauthorized),
	}
	if len(code) == 1 {
		errorObject.Code = code[0]
	}
	if len(code) > 1 {
		panic(errors.New("code cannot be more than one value"))
	}
	return errorObject
}
