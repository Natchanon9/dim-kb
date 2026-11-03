---
title: "Goroutine Leaks"
description: "Goroutines that never exit, and how to give them a way out"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

A leak is a goroutine that blocks forever: waiting on a channel nobody will use, or looping with no exit. Each one holds memory and whatever it references.

## Give every goroutine an exit

```go title="leak.go"
func counter(ctx context.Context, out chan<- int) {
	for i := 0; ; i++ {
		select {
		case out <- i:
		case <-ctx.Done(): // without this, the goroutine blocks forever
			return
		}
	}
}

ctx, cancel := context.WithCancel(context.Background())
out := make(chan int)
go counter(ctx, out)
fmt.Println(<-out, <-out) // 0 1
cancel() // counter returns
```

## Common causes

- Sending on an unbuffered channel after the receiver gave up (timeouts!). Fix: buffer of 1, or `select` with `ctx.Done()`.
- Ranging over a channel that is never closed.
- Forgetting to call `cancel()` or stop a `time.Ticker`.

Detect them with `runtime.NumGoroutine()` in tests, the pprof goroutine profile, or `go.uber.org/goleak`.

:::caution[Pitfalls]
- A slowly rising goroutine count in production metrics is the classic symptom.
:::
