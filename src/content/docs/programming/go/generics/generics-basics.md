---
title: "Generics Basics"
description: "Write one function or type that works for many types (Go 1.18+)"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Generics (Go 1.18+) let a function or type take **type parameters** in square brackets. You write the algorithm once and the compiler checks it for each type used.

## Example

```go title="map.go"
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) })
// ["2" "4" "6"]
```

## Before generics

You either copied the function per type, or used `any` and type assertions, losing compile-time checks. Generics keep both reuse and type safety.

:::caution[Pitfalls]
- Methods can't have their own type parameters; only the type can.
:::
