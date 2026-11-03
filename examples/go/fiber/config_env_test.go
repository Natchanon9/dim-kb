package gofiber

import (
	"os"
	"strconv"
	"testing"
	"time"
)

type Config struct {
	Port        int
	DatabaseURL string
	Timeout     time.Duration
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

// LoadConfig reads everything once at startup and fails fast on bad values.
func LoadConfig() (Config, error) {
	port, err := strconv.Atoi(getenv("PORT", "8080"))
	if err != nil {
		return Config{}, err
	}
	timeout, err := time.ParseDuration(getenv("TIMEOUT", "5s"))
	if err != nil {
		return Config{}, err
	}
	return Config{Port: port, DatabaseURL: os.Getenv("DATABASE_URL"), Timeout: timeout}, nil
}

func TestLoadConfig(t *testing.T) {
	t.Setenv("PORT", "9000")
	t.Setenv("DATABASE_URL", "postgres://localhost/app")
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != 9000 || cfg.Timeout != 5*time.Second || cfg.DatabaseURL == "" {
		t.Errorf("cfg = %+v", cfg)
	}

	t.Setenv("PORT", "abc")
	if _, err := LoadConfig(); err == nil {
		t.Error("expected an error for a bad PORT")
	}
}
