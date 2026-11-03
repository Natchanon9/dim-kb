---
title: "sync Package"
description: "The synchronisation primitives at a glance"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`sync` holds the low-level tools for coordinating goroutines. Most have their own page under Goroutines and Concurrency Patterns; this is the map.

## The toolbox

| Type | Use |
| --- | --- |
| `Mutex`, `RWMutex` | Protect shared state |
| `WaitGroup` | Wait for goroutines |
| `Once`, `OnceFunc`, `OnceValue(s)` | Run once, lazily |
| `Pool` | Reuse temporary objects |
| `Map` | Concurrent map for specific patterns |
| `Cond` | Wait for a condition (rare; channels are usually clearer) |

## sync.Map

```go title="syncmap.go"
var m sync.Map
m.Store("a", 1)
v, ok := m.Load("a")                    // 1 true
actual, loaded := m.LoadOrStore("a", 2) // 1 true (keeps existing)
```

It is faster than `map` + `Mutex` only when keys are written once and read many times, or goroutines touch disjoint keys. Otherwise use a regular map with a mutex: it's typed and simpler.

:::caution[Pitfalls]
- None of these may be copied after first use; pass pointers and run `go vet`.
:::
