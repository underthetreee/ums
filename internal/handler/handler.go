package handler

import (
	"github.com/go-chi/chi"
	"github.com/underthetreee/ums/internal/service"
)

func NewHandler(service service.User) *chi.Mux {
	r := chi.NewRouter()

	userHandler := NewUserHandler(service)
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
	})
	return r
}
