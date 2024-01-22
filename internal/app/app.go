package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/underthetreee/ums/internal/config"
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

	return nil
}
