package app

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/underthetreee/ums/internal/config"
	"github.com/underthetreee/ums/internal/server"
)

func Run() error {
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("failed to init config: %w", err)
	}

	db, err := sqlx.Connect("postgres", cfg.Postgres.URI)
	if err != nil {
		return err
	}
	defer db.Close()

	srv := server.NewServer(cfg)
	log.Printf("server is listening on :%s", cfg.HTTP.Port)
	if err := srv.Run(); err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}
