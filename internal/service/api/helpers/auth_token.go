package helpers

import (
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/config"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers/jwt"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"github.com/apodeixis/backend/resources"
	"github.com/pkg/errors"
)

func CreateAuthToken(user *data.User, r *http.Request, w http.ResponseWriter) (*resources.AuthToken, error) {
	accessExpiresAt := time.Now().Add(ctx.JwtConfig(r).AccessLife).Unix()
	claims := &jwt.Claims{
		OwnerId:   user.ID,
		ExpiresAt: accessExpiresAt,
	}
	access, err := jwt.CreateToken(claims, ctx.JwtConfig(r).Secret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create access token")
	}

	refreshExpiresAt := time.Now().Add(ctx.JwtConfig(r).RefreshLife).UTC()
	claims = &jwt.Claims{
		OwnerId:   user.ID,
		ExpiresAt: refreshExpiresAt.Unix(),
	}
	refresh, err := jwt.CreateToken(claims, ctx.JwtConfig(r).Secret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create refresh token")
	}

	err = ctx.RefreshTokensQ(r).New().FilterByUserID(user.ID).Delete()
	if err != nil && errors.Cause(err) != postgres.ErrNoSuchRefreshToken {
		return nil, err
	}
	refreshToken, err := ctx.RefreshTokensQ(r).New().Create(
		data.RefreshToken{
			UserID:    user.ID,
			Token:     refresh,
			ValidTill: refreshExpiresAt,
		},
	)
	if err != nil {
		return nil, err
	}
	setRefreshTokenCookie(refreshToken, w, ctx.RefreshCookieConfig(r))
	return responses.ConvertToAuthTokenResource(access, accessExpiresAt), nil
}

func setRefreshTokenCookie(refreshToken *data.RefreshToken, w http.ResponseWriter, cfg *config.RefreshCookieConfig) {
	maxAge := refreshToken.ValidTill.UTC().Sub(time.Now().UTC()).Seconds()
	cookie := &http.Cookie{
		Name:     cfg.Name,
		Value:    refreshToken.Token,
		Path:     cfg.Path,
		HttpOnly: cfg.HttpOnly,
		MaxAge:   int(maxAge),
		Secure:   cfg.Secure,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}
func RefreshAuthToken(refreshToken *data.RefreshToken, r *http.Request, w http.ResponseWriter) (*resources.AuthToken, error) {
	expiresAt := time.Now().Add(ctx.JwtConfig(r).AccessLife).Unix()
	claims := &jwt.Claims{
		OwnerId:   refreshToken.UserID,
		ExpiresAt: expiresAt,
	}
	access, err := jwt.CreateToken(claims, ctx.JwtConfig(r).Secret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create access token")
	}
	refreshToken, err = rotateRefreshToken(refreshToken, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to rotate refresh token")
	}
	setRefreshTokenCookie(refreshToken, w, ctx.RefreshCookieConfig(r))
	return responses.ConvertToAuthTokenResource(access, expiresAt), nil
}

func rotateRefreshToken(refreshToken *data.RefreshToken, r *http.Request) (*data.RefreshToken, error) {
	claims := &jwt.Claims{
		OwnerId:   refreshToken.UserID,
		ExpiresAt: refreshToken.ValidTill.Unix(),
	}
	refresh, err := jwt.CreateToken(claims, ctx.JwtConfig(r).Secret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create refresh token")
	}
	refreshToken.Token = refresh
	return ctx.RefreshTokensQ(r).New().FilterByID(refreshToken.ID).Update(*refreshToken)
}

func InvalidateRefreshTokenCookie(w http.ResponseWriter, cfg *config.RefreshCookieConfig) {
	cookie := &http.Cookie{
		Name:     cfg.Name,
		Value:    "",
		Path:     cfg.Path,
		HttpOnly: cfg.HttpOnly,
		MaxAge:   0,
		Secure:   cfg.Secure,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}
