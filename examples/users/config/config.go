package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds client configuration
type Config struct {
	SiteID string `envconfig:"SITE_ID" required:"true"`
	ApiKey string `envconfig:"API_KEY" required:"true"`
}

// Read reads config from .env file
func Read() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("godotenv.Load failed: %w", err)
	}

	var cnf Config
	err = envconfig.Process("", &cnf) // no prefix for env variables
	if err != nil {
		return nil, fmt.Errorf("godotenv.Load failed: %w", err)
	}

	return &cnf, nil
}
