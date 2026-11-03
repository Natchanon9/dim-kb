---
title: "log/slog"
description: "Structured logging in the standard library (Go 1.21+)"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`log/slog` writes structured logs: a message plus key-value attributes, as text or JSON. Log systems can then filter by field.

## Example

```go title="logger.go"
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

logger.Info("order created", "id", 42, "total", 99.5)
// {"time":"...","level":"INFO","msg":"order created","id":42,"total":99.5}

logger.With("request_id", "r1").Warn("slow query", slog.Duration("took", 2*time.Second))
// {...,"level":"WARN","msg":"slow query","request_id":"r1","took":2000000000}

logger.Debug("hidden") // below the level: not written
```

## Habits

- One logger built in `main`, passed down or set with `slog.SetDefault`.
- `logger.With(...)` attaches request-scoped fields once.
- Use `InfoContext(ctx, ...)` so handlers can add trace IDs from the context.
- JSON handler in production, text handler locally.

:::caution[Pitfalls]
- Odd numbers of key-value arguments produce a `!BADKEY` field; `go vet` catches it.
- Never log passwords, tokens or full card numbers.
:::
