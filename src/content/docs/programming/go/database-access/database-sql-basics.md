---
title: "database/sql Basics"
description: "One API for SQL databases, plus a driver per database"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`database/sql` is a generic interface over SQL databases. You add a **driver** for your database, and the same code works with PostgreSQL, MySQL or SQLite.

## Setup

```go title="db.go"
import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" driver
)

db, err := sql.Open("pgx", "postgres://app:secret@localhost:5432/app")
```

| Database | Driver import | Name |
| --- | --- | --- |
| PostgreSQL | `github.com/jackc/pgx/v5/stdlib` | `pgx` |
| MySQL | `github.com/go-sql-driver/mysql` | `mysql` |
| SQLite (pure Go) | `modernc.org/sqlite` | `sqlite` |

## Key types

- `*sql.DB`: a **pool** of connections, safe for concurrent use. Create one per database and share it.
- `*sql.Rows`, `*sql.Row`: query results.
- `*sql.Tx`: a transaction.
- `*sql.Stmt`: a prepared statement.

:::caution[Pitfalls]
- Don't open a `*sql.DB` per request; open once at startup and pass it in.
- Placeholders differ by driver: `$1` for Postgres, `?` for MySQL/SQLite.
:::
