---
title: "Mutex"
description: "Protect shared state with sync.Mutex and RWMutex"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A mutex lets one goroutine at a time into a critical section. Put the mutex next to the data it protects.

## Example

```go title="counter.go"
type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key]++
}
// 100 goroutines calling Inc("hits") → 100
```

## RWMutex

`sync.RWMutex` allows many readers or one writer. Use `RLock`/`RUnlock` for reads. It helps only when reads vastly outnumber writes; otherwise a plain `Mutex` is faster.

:::caution[Pitfalls]
- Never copy a struct holding a mutex; use pointer receivers.
- `Lock` is not reentrant: locking twice in the same goroutine deadlocks.
- Keep critical sections short; don't do I/O while holding a lock.
:::
