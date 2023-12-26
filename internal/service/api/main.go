package api

import (
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/backend/internal/config"
	"github.com/apodeixis/backend/internal/data/postgres"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/handlers"
	"github.com/apodeixis/backend/internal/service/api/middleware"
)

type API interface {
	Start() error
}

type api struct {
	router   chi.Router
	listener net.Listener
	log      *logan.Entry
}

func (a *api) Start() error {
	a.log.Info("Api started on ", a.listener.Addr().String())
	return http.Serve(a.listener, a.router)
}

func NewAPI(cfg config.Config) API {
	return &api{
		router:   newRouter(cfg),
		listener: cfg.Listener(),
		log:      cfg.Log(),
	}
}

func newRouter(cfg config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			ctx.SetLog(cfg.Log()),

			ctx.SetUsersQ(postgres.NewUsersQ(cfg.DB())),
			ctx.SetPostsQ(postgres.NewPostsQ(cfg.DB())),
			ctx.SetStarredPostsQ(postgres.NewStarredPostsQ(cfg.DB())),
			ctx.SetRefreshTokensQ(postgres.NewRefreshTokensQ(cfg.DB())),
			ctx.SetOAuth2StatesQ(postgres.NewOAuth2StatesQ(cfg.DB())),
			ctx.SetEmailVerificationTokensQ(postgres.NewEmailVerificationTokensQ(cfg.DB())),
			ctx.SetPasswordRecoveryTokensQ(postgres.NewPasswordRecoveryTokensQ(cfg.DB())),

			ctx.SetJwtConfig(cfg.JwtConfig()),
			ctx.SetOAuth2GoogleConfig(cfg.OAuth2GoogleConfig()),
			ctx.SetOAuth2GoogleStateConfig(cfg.OAuth2GoogleStateConfig()),
			ctx.SetRefreshCookieConfig(cfg.RefreshCookieConfig()),
			ctx.SetEmailVerificationConfig(cfg.EmailVerificationConfig()),
			ctx.SetPasswordRecoveryConfig(cfg.PasswordRecoveryConfig()),
			ctx.SetWebConfig(cfg.WebConfig()),
			ctx.SetNotificator(cfg.Notificator()),
		),
	)
	r.Route("/", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.With(middleware.Jwt()).Post("/logout", handlers.Logout)

			r.Route("/sign-up", func(r chi.Router) {
				r.Post("/", handlers.SignUp)
				r.Post("/email", handlers.SignUpEmail)
				r.Patch("/callback", handlers.SignUpCallback)
			})
			r.Post("/login", handlers.Login)
			r.Patch("/refresh", handlers.Refresh)

			r.Route("/oauth2/google", func(r chi.Router) {
				r.Post("/", handlers.OAuth2Google)
				r.Patch("/callback", handlers.OAuth2GoogleCallback)
			})

			r.Route("/recover/password", func(r chi.Router) {
				r.Patch("/", handlers.RecoverPassword)
				r.Post("/email", handlers.RecoverPasswordEmail)
			})
		})
		r.Route("/posts", func(r chi.Router) {
			r.With(middleware.Jwt()).Post("/", handlers.CreatePost)

			r.Get("/{id}", handlers.GetPost)
			r.Get("/headers/confirmed", handlers.GetConfirmedPostsHeaders)
			r.Get("/amount", handlers.GetPostsAmount)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.Jwt())
			r.Patch("/", handlers.EditUser)
			r.Route("/posts", func(r chi.Router) {
				r.Get("/headers", handlers.GetUserPostsHeaders)
				r.Route("/starred", func(r chi.Router) {
					r.Post("/", handlers.CreateStarredPost)
					r.Delete("/{id}", handlers.DeletePostFromStarred)
					r.Get("/headers", handlers.GetUserStarredPostsHeaders)
				})
			})
		})
		r.Get("/users/{id}", handlers.GetUser)
	})
	return r
}
