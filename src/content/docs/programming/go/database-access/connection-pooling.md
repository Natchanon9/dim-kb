---
title: "Connection Pooling"
description: "Size and tune the *sql.DB pool"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`*sql.DB` keeps a pool of connections. Defaults allow unlimited open connections, which can overwhelm the database under load. Set limits explicitly.

## Example

```go title="pool.go"
db.SetMaxOpenConns(10)                  // hard cap
db.SetMaxIdleConns(5)                   // kept warm between requests
db.SetConnMaxLifetime(30 * time.Minute) // recycle long-lived connections
db.SetConnMaxIdleTime(5 * time.Minute)

fmt.Println(db.Stats().MaxOpenConnections, db.Stats().OpenConnections) // 10 0
```

## Sizing

- Total across all instances must stay below the database's `max_connections` (Postgres default 100).
- 10–25 per instance is a common start; watch `db.Stats()` (`WaitCount`, `WaitDuration`) and tune.
- Serverless platforms with many instances often need a proxy like PgBouncer.

:::caution[Pitfalls]
- Unclosed `rows` and long transactions hold connections; the symptom is requests hanging on `WaitCount`.
:::
