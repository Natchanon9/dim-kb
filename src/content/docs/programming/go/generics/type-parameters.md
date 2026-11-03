---
title: "Type Parameters"
description: "Declare [T Constraint] and instantiate with concrete types"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A type parameter list `[T Constraint]` comes after the function name. Each parameter has a constraint saying what the type must support.

## Example

```go title="larger.go"
func Larger[T cmp.Ordered](a, b T) T { // T must support < and >
	if a > b {
		return a
	}
	return b
}

Larger(3, 7)             // 7, T inferred as int
Larger("a", "b")         // b
Larger[float64](1, 2.5)  // 2.5, T given explicitly
```

Several parameters: `[K comparable, V any]`. Parameters sharing a constraint: `[T, U any]`.

:::caution[Pitfalls]
- Both arguments must have the **same** `T`: `Larger(1, 2.5)` infers `float64` from the untyped constants, but `Larger(x, y)` with `int` and `float64` variables fails.
:::
