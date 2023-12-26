package handlers

import (
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/service/api/responses"

	"github.com/apodeixis/backend/internal/config"
	"github.com/apodeixis/backend/internal/service/api/helpers"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	hoauth2 "github.com/apodeixis/backend/internal/service/api/helpers/oauth2"
	"github.com/apodeixis/backend/internal/service/api/helpers/oauth2/google"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"golang.org/x/oauth2"
)

func OAuth2GoogleCallback(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewOAuth2GoogleCallbackRequest(r)
	if err != nil {
		ctx.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	state := request.Data.Attributes.State
	oAuth2State, err := ctx.OAuth2StatesQ(r).New().FilterByState(state).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if oAuth2State == nil {
		ctx.Log(r).Error("not found state")
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundState))
		return
	}
	if err := verifyOAuth2State(oAuth2State, ctx.OAuth2GoogleStateConfig(r)); err != nil {
		ctx.Log(r).WithError(err).Error("oauth2 state is invalid")
		ape.RenderErr(w, problems.Forbidden(types.ErrInvalidOAuth2State))
		return
	}
	user, err := getUserFromGoogle(request.Data.Attributes.Code, ctx.OAuth2GoogleConfig(r))
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to get user from google")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	savedUser, err := ctx.UsersQ(r).New().FilterByEmail(user.Email).Get()
	if err != nil {
		ctx.Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if savedUser != nil && !savedUser.OAuth2User {
		ctx.Log(r).Error("not oauth2 user tries to use oauth2")
		ape.RenderErr(w, problems.UnprocessableEntity(types.ErrUnprocessableEntityNotOAuth2User))
		return
	}
	_ = ctx.UsersQ(r).New().Transaction(func() error {
		err := ctx.OAuth2StatesQ(r).New().FilterByID(oAuth2State.ID).Delete()
		if err != nil {
			ctx.Log(r).Error(err)
			ape.RenderErr(w, problems.InternalError())
			return err
		}
		var createdNewUser bool
		if savedUser == nil {
			user.AuthorID = helpers.GenerateID()
			provider := types.GoogleOAuth2Provider
			user.OAuth2Provider = &provider
			user, err = ctx.UsersQ(r).New().CreateUserFromOAuth2(*user)
			if err != nil {
				ctx.Log(r).Error(err)
				ape.RenderErr(w, problems.InternalError())
				return err
			}
			createdNewUser = true
		} else {
			user = savedUser
		}
		authToken, err := helpers.CreateAuthToken(user, r, w)
		if err != nil {
			ctx.Log(r).WithError(err).Error("failed to create auth token")
			ape.RenderErr(w, problems.InternalError())
			return err
		}
		response := responses.ComposeUserWithAuthToken(user, authToken)
		if createdNewUser {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		ape.Render(w, response)
		return nil
	})
}

func verifyOAuth2State(oAuth2State *data.OAuth2State, cfg *config.OAuth2GoogleStateConfig) error {
	if oAuth2State.ValidTill.UTC().Before(time.Now().UTC()) {
		return errors.New("oauth2 state is expired")
	}
	err := helpers.VerifyToken(oAuth2State.State, cfg.StateSecret)
	return errors.Wrap(err, "invalid token")
}

func getUserFromGoogle(code string, oauthCfg *oauth2.Config) (*data.User, error) {
	token, err := hoauth2.GetUserToken(code, oauthCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user token")
	}
	userInfo, err := google.GetUserInfo(token, oauthCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}
	return google.ConvertUserInfoToUser(userInfo)
}
