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
	"github.com/underthetreee/ums/internal/handler"
	"github.com/underthetreee/ums/internal/repository"
	"github.com/underthetreee/ums/internal/server"
	"github.com/underthetreee/ums/internal/service"
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

	userService := service.NewUserService(
		repository.NewUserRepo(db),
	)

	handler := handler.NewHandler(userService)
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
