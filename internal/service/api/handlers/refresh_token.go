package handlers

import (
	"net/http"

	"github.com/apodeixis/backend/internal/types"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/helpers/jwt"
	"github.com/apodeixis/backend/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/service/api/requests"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRefreshTokenRequest(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	refreshToken, err := ctx.RefreshTokensQ(r).New().FilterByToken(request.Token).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if refreshToken == nil {
		ctx.Log(r).Error("not found refresh token")
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundRefreshToken))
		return
	}
	_, err = jwt.ExtractClaims(refreshToken.Token, ctx.JwtConfig(r).Secret)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to extract claims")
		ape.RenderErr(w, problems.Unauthorized(types.ErrInvalidJWTToken))
		return
	}
	authToken, err := helpers.RefreshAuthToken(refreshToken, r, w)
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeAuthToken(authToken)
	w.WriteHeader(http.StatusOK)
	ape.Render(w, response)
}
