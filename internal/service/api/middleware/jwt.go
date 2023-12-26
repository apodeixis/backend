package middleware

import (
	"net/http"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Jwt() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := requests.ExtractClaimsFromAuthHeader(r)
			if err != nil {
				ctx.Log(r).Errorf("unauthorized: %s", err)
				ape.RenderErr(w, problems.Unauthorized())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
