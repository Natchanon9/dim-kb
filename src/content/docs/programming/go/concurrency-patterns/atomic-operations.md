---
title: "Atomic Operations"
description: "Lock-free counters and flags with sync/atomic"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`sync/atomic` updates a single value without a lock. It fits counters, flags and swapping a pointer to immutable config.

## Example

```go title="atomic.go"
var hits atomic.Int64
// 1000 goroutines call hits.Add(1)
fmt.Println(hits.Load()) // 1000

var state atomic.Int32
state.CompareAndSwap(0, 1) // true: 0 → 1
state.CompareAndSwap(0, 1) // false: already 1
```

Typed atomics (`atomic.Int64`, `atomic.Bool`, `atomic.Pointer[T]`, Go 1.19+) are safer than the old `atomic.AddInt64(&x, 1)` functions.

## Atomic or mutex?

One independent value → atomic. Several values that must change together → mutex.

:::caution[Pitfalls]
- Mixing atomic and plain access to the same variable is still a data race.
:::
