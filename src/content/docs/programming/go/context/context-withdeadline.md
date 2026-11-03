---
title: "context.WithDeadline"
description: "Cancel at a fixed point in time"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`WithDeadline(parent, t)` cancels at time `t`. `WithTimeout(parent, d)` is just `WithDeadline(parent, time.Now().Add(d))`.

## Example

```go title="deadline.go"
deadline := time.Now().Add(50 * time.Millisecond)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
d, ok := ctx.Deadline() // deadline, true

child, cancelChild := context.WithTimeout(ctx, time.Hour)
defer cancelChild()
cd, _ := child.Deadline() // still the parent's deadline
```

## When a fixed time fits better

When a whole job must finish by one moment (end of a batch window, an SLA for a multi-step request), compute the deadline once and pass the same context to every step.

:::caution[Pitfalls]
- Check remaining time with `time.Until(d)` before starting work that can't finish in time.
:::
