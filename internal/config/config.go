package config

import "os"

type Config struct {
	DatabaseURL string
}

func New() *Config {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
