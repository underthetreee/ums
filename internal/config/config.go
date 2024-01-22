package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		URI string `envconfig:"POSTGRES_URI"`
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
