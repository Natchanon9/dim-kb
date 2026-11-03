---
title: "Context Basics"
description: "Carry cancellation, deadlines and request values across calls"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A `context.Context` carries a cancellation signal, an optional deadline, and request-scoped values. Every blocking function in a server should accept one.

## Example

```go title="ctx.go"
func slowOp(ctx context.Context) error {
	select {
	case <-time.After(time.Second): // the real work
		return nil
	case <-ctx.Done(): // cancelled or timed out
		return ctx.Err()
	}
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
defer cancel()
fmt.Println(slowOp(ctx)) // context deadline exceeded
```

## The interface

| Method | Returns |
| --- | --- |
| `Done()` | A channel closed when the context ends |
| `Err()` | `nil`, `context.Canceled` or `context.DeadlineExceeded` |
| `Deadline()` | The deadline, if any |
| `Value(key)` | A request-scoped value |

Roots: `context.Background()` in `main`/tests; `context.TODO()` as a placeholder. In HTTP handlers use `r.Context()`.

:::caution[Pitfalls]
- Pass `ctx` as the **first** parameter; never store it in a struct.
- Never pass a `nil` context; use `context.TODO()`.
:::
