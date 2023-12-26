package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

// UnprocessableEntity composes JSON:API error object with
// optional application-specific error code.
func UnprocessableEntity(code ...string) *jsonapi.ErrorObject {
	errorObject := &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusUnprocessableEntity),
		Status: fmt.Sprintf("%d", http.StatusUnprocessableEntity),
	}
	if len(code) == 1 {
		errorObject.Code = code[0]
	}
	if len(code) > 1 {
		panic(errors.New("code cannot be more than one value"))
	}
	return errorObject
}
