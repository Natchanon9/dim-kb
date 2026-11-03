---
title: "Config & ENV"
description: "Load configuration from environment variables at startup"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Read config once at startup into a struct, validate it, and pass it down. Environment variables are the standard for containers (12-factor).

## Example

```go title="config.go"
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
```

Test with `t.Setenv("PORT", "9000")`.

## .env files locally

`github.com/joho/godotenv` loads a `.env` file into the environment for local development: `godotenv.Load()` at the top of `main`. Libraries like `caarlos0/env` or Viper map env vars to struct tags.

:::caution[Pitfalls]
- Commit `.env.example`, never `.env` with real secrets.
- Fail at startup on missing required config, not on the first request.
:::
