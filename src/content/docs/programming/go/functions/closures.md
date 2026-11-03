---
title: "Closures"
description: "Functions that capture and keep outer variables"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

A closure is a function that uses variables from the scope where it was created. Those variables live as long as the closure does.

## Counter

```go title="counter.go"
func counter() func() int {
	n := 0
	return func() int {
		n++ // n lives on between calls
		return n
	}
}

next := counter()
next(); next()
other := counter()          // its own n
fmt.Println(next(), other()) // 3 1
```

Closures capture variables **by reference**, so each `counter()` call creates separate state.

## Real uses

Middleware that wraps a handler, functional options, memoization, and generators.

:::caution[Pitfalls]
- Captured variables can be shared with goroutines; protect them with a mutex if more than one goroutine writes.
- Before Go 1.22, closures in a loop all captured the same loop variable.
:::
