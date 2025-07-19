package config

import (
	"github.com/caarlos0/env/v6"
)

func NewConfig() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

type Config struct {
	HTTPAddress          string `env:"RUN_ADDRESS" envDefault:":8080"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS" envDefault:"http://localhost:8081"`
	DSN                  string `env:"DATABASE_URI" envDefault:"user=postgres password=postgres dbname=praktikum host=localhost port=5432 sslmode=disable"`
	HashSecret           string `env:"HASH_SECRET" envDefault:"my_secret"`
	AuthUserCookieName   string `env:"AUTH_USER_COOKIE_NAME" envDefault:"Authorization"`
}
