package handlers

import (
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/service/api/helpers"

	"github.com/apodeixis/backend/internal/service/api/ctx"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/backend/internal/data"
)

func OAuth2Google(w http.ResponseWriter, r *http.Request) {
	state, err := helpers.GenerateToken(ctx.OAuth2GoogleStateConfig(r).StateSecret)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to generate state string")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	validTill := time.Now().Add(ctx.OAuth2GoogleStateConfig(r).StateLife).UTC()
	_, err = ctx.OAuth2StatesQ(r).New().Create(data.OAuth2State{
		State:     state,
		ValidTill: &validTill,
	})
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	url := ctx.OAuth2GoogleConfig(r).AuthCodeURL(state)
	w.WriteHeader(http.StatusOK)
	response := responses.ComposeOAuth2Google(url)
	ape.Render(w, response)
}
