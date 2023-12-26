package problems

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/google/jsonapi"
)

// Forbidden composes JSON:API error object with
// optional application-specific error code.
func Forbidden(code ...string) *jsonapi.ErrorObject {
	errorObject := &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusForbidden),
		Status: fmt.Sprintf("%d", http.StatusForbidden),
	}
	if len(code) == 1 {
		errorObject.Code = code[0]
	}
	if len(code) > 1 {
		panic(errors.New("code cannot be more than one value"))
	}
	return errorObject
}
