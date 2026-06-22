package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port: os.Getenv("PORT"),
	}

	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT is required")
	}

	return cfg, nil
}
