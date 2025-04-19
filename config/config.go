package config

import (
	"flag"
	"os"
)

type Config struct {
	HTTPAddress          string
	AccrualSystemAddress string
	FileStoragePath      string
	DSN                  string
	HashSecret           string
	AuthUserCookieName   string
}

var cfg Config

func Init() Config {
	flag.StringVar(&cfg.HTTPAddress, "a", ":8080", "HTTP run address")
	flag.StringVar(&cfg.AccrualSystemAddress, "r", "localhost:8081", "accrual run url")
	flag.StringVar(&cfg.DSN, "d", "user=postgres password=postgres dbname=praktikum host=localhost port=5432 sslmode=disable", "db connection")
	flag.StringVar(&cfg.HashSecret, "h", "my_secret", "hash secret")
	flag.StringVar(&cfg.AuthUserCookieName, "c", "auth_cookie", "auth cookie name")

	flag.Parse()

	if httpAddress := os.Getenv("RUN_ADDRESS"); httpAddress != "" {
		cfg.HTTPAddress = httpAddress
	}

	if accrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); accrualSystemAddress != "" {
		cfg.AccrualSystemAddress = accrualSystemAddress
	}

	if dsn := os.Getenv("DSN"); dsn != "" {
		cfg.DSN = dsn
	}

	if hashSecret := os.Getenv("HASH_SECRET"); hashSecret != "" {
		cfg.HashSecret = hashSecret
	}

	if authUserCookieName := os.Getenv("AUTH_USER_COOKIE_NAME"); authUserCookieName != "" {
		cfg.AuthUserCookieName = authUserCookieName
	}

	return cfg
}
