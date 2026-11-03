---
title: "sqlx Package"
description: "Scan rows into structs and expand IN clauses"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`github.com/jmoiron/sqlx` extends `database/sql` with struct scanning by `db` tag, named parameters and `IN` expansion, while keeping plain SQL.

## Scanning into structs

```go title="users.go"
type userRow struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
}

db := sqlx.MustConnect("pgx", dsn)
var users []userRow
err := db.SelectContext(ctx, &users, `SELECT id, email FROM users WHERE active`)

var u userRow
err = db.GetContext(ctx, &u, `SELECT id, email FROM users WHERE id = $1`, id)
```

## IN clauses

```go title="in.go"
q, args, _ := sqlx.In(`SELECT id FROM users WHERE id IN (?)`, []int{1, 2, 3})
q = sqlx.Rebind(sqlx.DOLLAR, q) // or db.Rebind(q)
// SELECT id FROM users WHERE id IN ($1, $2, $3) [1 2 3]
```

## Where it sits

| Tool | Style |
| --- | --- |
| `database/sql` | Manual `Scan` |
| `sqlx` | SQL + struct mapping |
| `sqlc` | Generates Go from SQL files |
| GORM | Full ORM |

:::caution[Pitfalls]
- `Select` loads all rows into memory; use `Queryx` + `StructScan` for big result sets.
:::
