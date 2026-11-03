---
title: "Iterators"
description: "iter.Seq, iter.Seq2 and iter.Pull"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

The `iter` package (Go 1.23+) names the iterator types: `iter.Seq[V]` yields values, `iter.Seq2[K, V]` yields pairs. Standard packages return these instead of building slices.

## Seq2 and Pull

```go title="iter.go"
func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

for i, s := range Enumerate([]string{"a", "b"}) { ... } // 0 a, 1 b

next, stop := iter.Pull(Countdown(3)) // pull values one at a time
defer stop()
v1, _ := next() // 3
v2, _ := next() // 2
```

## Why iterators

- Lazy: values are produced only as consumed, so infinite or huge sequences are fine.
- No intermediate slices.
- One standard shape across libraries (`maps.Keys`, `slices.Values`, `strings.Lines` in 1.24).

:::caution[Pitfalls]
- Always `defer stop()` after `iter.Pull`, or the iterator's resources leak.
:::
