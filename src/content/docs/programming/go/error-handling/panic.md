---
title: "panic"
description: "Stop the program for bugs and impossible states"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`panic` stops normal execution, runs deferred calls up the stack, and crashes the program with a stack trace unless something recovers.

## Example

```go title="panic.go"
func mustPositive(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("mustPositive: got %d", n))
	}
	return n
}
```

The runtime also panics on its own: nil dereference, index out of range, nil map write, closing a closed channel.

## When to panic

- Programmer errors that should never happen at runtime.
- Startup failures where continuing makes no sense (`regexp.MustCompile` on a constant, missing required config in `main`).

Everything a caller could reasonably handle (bad input, network, files) is an `error`, not a panic.

:::caution[Pitfalls]
- A panic in any goroutine crashes the whole program, not just that goroutine.
- Libraries should not panic across their API.
:::
