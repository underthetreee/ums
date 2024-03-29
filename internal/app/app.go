package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/underthetreee/ums/internal/config"
	handler "github.com/underthetreee/ums/internal/http"
	"github.com/underthetreee/ums/internal/server"
)

func Run() error {
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("init config: %w", err)
	}

	db, err := sqlx.Connect("postgres", cfg.Postgres.URI)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}
	defer db.Close()

	seedUsersTable(db)

	handler := handler.NewHandler(db)
	srv := server.NewServer(cfg, handler)

	var (
		quitch = make(chan os.Signal, 1)
		errch  = make(chan error, 1)
	)

	go func() {
		if err := srv.Run(); err != http.ErrServerClosed {
			errch <- err
		}
	}()
	log.Printf("server is listening on :%s", cfg.HTTP.Port)

	signal.Notify(quitch, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quitch:
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			return fmt.Errorf("stop http server: %w", err)
		}
	case err := <-errch:
		return fmt.Errorf("start http server: %w", err)
	}
	return nil
}

func seedUsersTable(db *sqlx.DB) {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	db.MustExec(usersTable)
}
