package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/underthetreee/ums/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, router *chi.Mux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.HTTP.Port,
			Handler:      router,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
