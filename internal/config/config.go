package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
	}

	PostgresConfig struct {
		URI string `envconfig:"POSTGRES_URI"`
	}

	HTTPConfig struct {
		Port         string        `envconfig:"HTTP_PORT"`
		ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT"`
		WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT"`
	}
)

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
