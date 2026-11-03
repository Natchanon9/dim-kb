---
title: "context.WithTimeout"
description: "Cancel automatically after a duration"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`WithTimeout(parent, d)` cancels itself after `d`. It is the standard way to bound how long an outgoing call may take.

## Example

```go title="timeout.go"
ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
defer cancel()
err := slowOp(ctx)
errors.Is(err, context.DeadlineExceeded) // true
```

## Typical uses

```go
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()
row := db.QueryRowContext(ctx, "SELECT ...")
```

Same for `http.NewRequestWithContext`, gRPC calls and Redis clients.

:::caution[Pitfalls]
- The timeout only works if the called code watches `ctx.Done()`. Libraries do; your own loops must too.
- A child timeout can't outlive the parent's; the earlier deadline wins.
:::
