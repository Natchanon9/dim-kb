---
title: "Goroutine Patterns"
description: "Generators and bounded concurrency"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

A few shapes cover most goroutine use. The rule behind them: the goroutine that creates a channel owns it and closes it.

## Generator

```go title="gen.go"
func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range n {
			ch <- i
		}
	}()
	return ch
}

for v := range gen(3) { ... } // 0, 1, 2
```

## Bounded concurrency (semaphore)

```go title="sem.go"
sem := make(chan struct{}, 2) // at most 2 at once
for range 5 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{}        // acquire
		defer func() { <-sem }() // release
		doWork()
	}()
}
wg.Wait()
```

Other patterns (worker pools, pipelines, fan-in/fan-out, errgroup) are in [Channels](/dim-kb/programming/go/channels/channel-basics/) and [Concurrency Patterns](/dim-kb/programming/go/concurrency-patterns/concurrency-overview/).

:::caution[Pitfalls]
- Return receive-only channels (`<-chan T`) so callers can't close or send on them.
:::
