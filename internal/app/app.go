package app

import (
	"fmt"

	"github.com/underthetreee/ums/internal/config"
)

func Run() error {
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("failed to init config: %w", err)
	}
	fmt.Println(cfg)
	return nil
}
