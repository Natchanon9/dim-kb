---
title: "Connecting to Database"
description: "Open, ping and configure the pool at startup"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`sql.Open` only validates its arguments; the first real connection happens on first use. Ping at startup so a wrong DSN fails fast.

## Example

```go title="open.go"
func openDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}
```

## DSN

```text
postgres://user:password@host:5432/dbname?sslmode=disable
```

Read it from an environment variable (`DATABASE_URL`), never hard-code credentials. Locally, run Postgres with Docker:

```sh
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:17
```

:::caution[Pitfalls]
- `defer db.Close()` belongs in `main`, not in a function that returns the db.
:::
