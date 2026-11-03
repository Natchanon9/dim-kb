---
title: "Type Sets"
description: "What ~T means and how constraints define a set of types"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Every constraint describes a **type set**: the types that satisfy it. `int` means exactly `int`; `~int` means every type whose underlying type is `int`.

## Example

```go title="avg.go"
type Celsius float64

type Float interface {
	~float32 | ~float64 // ~ includes named types like Celsius
}

func Avg[T Float](xs ...T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total / T(len(xs))
}

Avg(Celsius(20), Celsius(30)) // 25
```

Without `~`, `Celsius` would not be allowed, even though it is a `float64` underneath.

## Combining

- `A | B`: union of the two sets.
- Lines in one interface intersect: `interface{ ~int | ~string; String() string }` = int- or string-based types that also have `String()`.

:::caution[Pitfalls]
- Prefer `~` in reusable constraints, or user-defined types like `type UserID int` won't fit.
:::
