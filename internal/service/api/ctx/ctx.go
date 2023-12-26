package ctx

import (
	"context"
	"net/http"

	"github.com/apodeixis/backend/internal/connectors/notificator"

	"github.com/apodeixis/backend/internal/config"
	"golang.org/x/oauth2"

	"github.com/apodeixis/backend/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logKey ctxKey = iota + 1

	oAuth2StatesQCtxKey
	usersQCtxKey
	postsQCtxKey
	starredPostsQCtxKey
	refreshTokesQCtxKey
	emailVerificationTokensQCtxKey
	passwordRecoveryTokensQCtxKey

	jwtConfigCtxKey
	oAuth2GoogleConfigCtxKey
	oAuth2GoogleStateConfigCtxKey
	refreshCookieConfigCtxKey
	emailVerificationConfigCtxKey
	passwordRecoveryConfigCtxKey
	webConfigCtxKey

	notificatorCtxKey
)

func SetLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logKey).(*logan.Entry)
}

func SetUsersQ(q data.Users) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersQCtxKey, q)
	}
}

func UsersQ(r *http.Request) data.Users {
	return r.Context().Value(usersQCtxKey).(data.Users)
}

func SetPostsQ(q data.Posts) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, postsQCtxKey, q)
	}
}

func PostsQ(r *http.Request) data.Posts {
	return r.Context().Value(postsQCtxKey).(data.Posts)
}

func SetStarredPostsQ(q data.StarredPosts) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, starredPostsQCtxKey, q)
	}
}

func StarredPostsQ(r *http.Request) data.StarredPosts {
	return r.Context().Value(starredPostsQCtxKey).(data.StarredPosts)
}

func SetRefreshTokensQ(q data.RefreshTokens) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, refreshTokesQCtxKey, q)
	}
}

func RefreshTokensQ(r *http.Request) data.RefreshTokens {
	return r.Context().Value(refreshTokesQCtxKey).(data.RefreshTokens)
}

func SetOAuth2StatesQ(q data.OAuth2States) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2StatesQCtxKey, q)
	}
}

func OAuth2StatesQ(r *http.Request) data.OAuth2States {
	return r.Context().Value(oAuth2StatesQCtxKey).(data.OAuth2States)
}

func SetEmailVerificationTokensQ(q data.EmailVerificationTokens) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, emailVerificationTokensQCtxKey, q)
	}
}

func EmailVerificationTokensQ(r *http.Request) data.EmailVerificationTokens {
	return r.Context().Value(emailVerificationTokensQCtxKey).(data.EmailVerificationTokens)
}

func SetPasswordRecoveryTokensQ(q data.PasswordRecoveryTokens) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, passwordRecoveryTokensQCtxKey, q)
	}
}

func PasswordRecoveryTokensQ(r *http.Request) data.PasswordRecoveryTokens {
	return r.Context().Value(passwordRecoveryTokensQCtxKey).(data.PasswordRecoveryTokens)
}

func SetJwtConfig(cfg *config.JwtConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtConfigCtxKey, cfg)
	}
}

func JwtConfig(r *http.Request) *config.JwtConfig {
	return r.Context().Value(jwtConfigCtxKey).(*config.JwtConfig)
}

func SetOAuth2GoogleConfig(cfg *oauth2.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2GoogleConfigCtxKey, cfg)
	}
}

func OAuth2GoogleConfig(r *http.Request) *oauth2.Config {
	return r.Context().Value(oAuth2GoogleConfigCtxKey).(*oauth2.Config)
}

func SetOAuth2GoogleStateConfig(cfg *config.OAuth2GoogleStateConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2GoogleStateConfigCtxKey, cfg)
	}
}

func OAuth2GoogleStateConfig(r *http.Request) *config.OAuth2GoogleStateConfig {
	return r.Context().Value(oAuth2GoogleStateConfigCtxKey).(*config.OAuth2GoogleStateConfig)
}

func SetRefreshCookieConfig(cfg *config.RefreshCookieConfig) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, refreshCookieConfigCtxKey, cfg)
	}
}

func RefreshCookieConfig(r *http.Request) *config.RefreshCookieConfig {
	return r.Context().Value(refreshCookieConfigCtxKey).(*config.RefreshCookieConfig)
}

func SetEmailVerificationConfig(cfg *config.EmailVerificationConfig) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, emailVerificationConfigCtxKey, cfg)
	}
}

func EmailVerificationConfig(r *http.Request) *config.EmailVerificationConfig {
	return r.Context().Value(emailVerificationConfigCtxKey).(*config.EmailVerificationConfig)
}

func SetPasswordRecoveryConfig(cfg *config.PasswordRecoveryConfig) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, passwordRecoveryConfigCtxKey, cfg)
	}
}

func PasswordRecoveryConfig(r *http.Request) *config.PasswordRecoveryConfig {
	return r.Context().Value(passwordRecoveryConfigCtxKey).(*config.PasswordRecoveryConfig)
}

func SetWebConfig(cfg *config.WebConfig) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, webConfigCtxKey, cfg)
	}
}

func WebConfig(r *http.Request) *config.WebConfig {
	return r.Context().Value(webConfigCtxKey).(*config.WebConfig)
}

func SetNotificator(notificator notificator.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, notificatorCtxKey, notificator)
	}
}

func Notificator(r *http.Request) notificator.Connector {
	return r.Context().Value(notificatorCtxKey).(notificator.Connector)
}
