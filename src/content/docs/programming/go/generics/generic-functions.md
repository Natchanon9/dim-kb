---
title: "Generic Functions"
description: "Filter, Reduce and other reusable helpers"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Generic functions shine for small algorithms over collections where only the element type changes.

## Example

```go title="filter.go"
func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

evens := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }) // [2 4]
total := Reduce(evens, 0, func(acc, n int) int { return acc + n })        // 6
```

Check the standard `slices` and `maps` packages first; many helpers (`slices.IndexFunc`, `slices.DeleteFunc`) already exist.

:::caution[Pitfalls]
- Chaining Filter/Map/Reduce is less idiomatic in Go than a plain `for` loop; use helpers where they make code clearer.
:::
