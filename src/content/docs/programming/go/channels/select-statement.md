---
title: "Select Statement"
description: "Wait on several channel operations at once"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`select` waits until one of its cases can proceed, then runs that case. If several are ready, it picks one at random.

## Timeout and non-blocking

```go title="select.go"
fast := make(chan string, 1)
fast <- "cache"
select {
case v := <-fast:
	fmt.Println(v) // cache
case <-time.After(time.Second):
	fmt.Println("timeout")
}

empty := make(chan int)
select {
case v := <-empty:
	fmt.Println(v)
default: // runs immediately when no case is ready
	fmt.Println("no value")
}
```

## Common shapes

- **Timeout:** a case on `time.After` or `ctx.Done()`.
- **Non-blocking send/receive:** add `default`.
- **Event loop:** `for { select { ... } }` with a quit case.

:::caution[Pitfalls]
- `time.After` inside a hot loop creates a timer each iteration; reuse a `time.Timer` or use a context.
- An empty `select {}` blocks forever.
:::
