package problems

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/google/jsonapi"
)

// InternalError composes JSON:API error object with
// optional application-specific error code.
func InternalError(code ...string) *jsonapi.ErrorObject {
	errorObject := &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
	}
	if len(code) == 1 {
		errorObject.Code = code[0]
	}
	if len(code) > 1 {
		panic(errors.New("code cannot be more than one value"))
	}
	return errorObject
}
