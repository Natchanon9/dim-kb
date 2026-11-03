---
title: "Context Propagation"
description: "Pass the same ctx through every layer of a request"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A request's context flows from the handler down through services and repositories to the database or HTTP client. When the client disconnects or the deadline passes, every layer stops.

## Example

```go title="layers.go"
func repoFind(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("repo: %w", err)
	}
	return db.QueryRowContext(ctx, "...").Err()
}

func serviceGet(ctx context.Context) error {
	if err := repoFind(ctx); err != nil {
		return fmt.Errorf("service: %w", err)
	}
	return nil
}

// cancelled request → service: repo: context canceled
```

## Rules

- Accept `ctx` first in every function that does I/O and pass it on unchanged, or derive a child (`WithTimeout`).
- Use the `...Context` variants: `QueryContext`, `NewRequestWithContext`.
- Work that must outlive the request (audit log, async email) uses `context.WithoutCancel(ctx)` (Go 1.21+) so values survive but cancellation doesn't.

:::caution[Pitfalls]
- Starting a goroutine with the request `ctx` and returning: the work gets cancelled as soon as the handler returns.
:::
