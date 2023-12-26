package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/types"

	"github.com/apodeixis/backend/internal/service/api/helpers"

	"github.com/pkg/errors"

	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewLogout(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = ctx.RefreshTokensQ(r).New().FilterByUserID(request.UserID).Delete()
	if errors.Is(err, postgres.ErrNoSuchRefreshToken) {
		log.Error(err)
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundRefreshToken))
		return
	}
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	helpers.InvalidateRefreshTokenCookie(w, ctx.RefreshCookieConfig(r))
	w.WriteHeader(http.StatusNoContent)
}
