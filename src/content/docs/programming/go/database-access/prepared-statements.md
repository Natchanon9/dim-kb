---
title: "Prepared Statements"
description: "Parse a query once and execute it many times"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A prepared statement is parsed and planned by the database once, then executed with different arguments. Useful for the same query in a tight loop.

## Example

```go title="insert.go"
stmt, err := db.PrepareContext(ctx, `INSERT INTO users (name) VALUES ($1)`)
if err != nil {
	return err
}
defer stmt.Close()
for _, n := range names {
	if _, err := stmt.ExecContext(ctx, n); err != nil {
		return err
	}
}
```

## Do you need it?

Plain `db.QueryContext(ctx, query, args...)` with placeholders is already safe from SQL injection, and drivers like pgx cache statements automatically. Prepare explicitly only for hot loops; for bulk inserts, multi-row `INSERT` or `COPY` is faster still.

:::caution[Pitfalls]
- A `*sql.Stmt` from `db` re-prepares on each new connection; inside a transaction use `tx.StmtContext`.
:::
