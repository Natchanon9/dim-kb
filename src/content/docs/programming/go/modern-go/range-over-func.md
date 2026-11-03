---
title: "range over func"
description: "Write your own iterators that work with for-range (Go 1.23+)"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Since Go 1.23, `for range` accepts a function. The function calls `yield` once per value; `yield` returns false when the loop stops early.

## Example

```go title="countdown.go"
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i > 0; i-- {
			if !yield(i) { // loop body hit break/return
				return
			}
		}
	}
}

for v := range Countdown(5) {
	if v == 2 {
		break
	}
	fmt.Println(v) // 5 4 3
}
```

## Accepted shapes

| Function type | Loop form |
| --- | --- |
| `func(yield func() bool)` | `for range f` |
| `func(yield func(V) bool)` | `for v := range f` |
| `func(yield func(K, V) bool)` | `for k, v := range f` |

:::caution[Pitfalls]
- Always check `yield`'s result; calling it again after it returned false panics.
- Requires `go 1.23` in `go.mod` (or a `//go:build go1.23` file).
:::
