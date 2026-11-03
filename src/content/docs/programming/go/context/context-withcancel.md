---
title: "context.WithCancel"
description: "A context you can stop by calling cancel()"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`WithCancel` returns a child context and a `cancel` function. Calling `cancel` closes `Done()` for that context and all its children.

## Example

```go title="cancel.go"
ctx, cancel := context.WithCancel(context.Background())
cancel()
<-ctx.Done()
fmt.Println(ctx.Err(), errors.Is(ctx.Err(), context.Canceled)) // context canceled true

ctx2, cancel2 := context.WithCancelCause(context.Background()) // Go 1.20+
cancel2(errors.New("user logged out"))
fmt.Println(context.Cause(ctx2)) // user logged out
```

Use it to stop background workers on shutdown, or to stop sibling goroutines once one has the answer.

:::caution[Pitfalls]
- Always call `cancel`, usually with `defer cancel()`, even if the work finishes normally; otherwise the context's resources leak until the parent ends (`go vet` warns).
- `cancel` is safe to call more than once.
:::
