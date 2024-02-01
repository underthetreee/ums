package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/underthetreee/ums/internal/auth"
	v1 "github.com/underthetreee/ums/internal/http/v1"
	"github.com/underthetreee/ums/internal/repository"
	"github.com/underthetreee/ums/internal/service"
)

func NewHandler(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userHandler := v1.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))

	r.Route("/v1/api", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)

		r.Route("/profile", func(r chi.Router) {
			r.Use(auth.JWTAuth)

			r.Get("/", userHandler.GetProfile)
			r.Put("/", userHandler.UpdateProfile)
		})
	})
	return r
}
