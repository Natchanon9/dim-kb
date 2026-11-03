---
title: "slices Package (Modern)"
description: "Iterator-based helpers added in Go 1.23"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go 1.23 connected `slices` to iterators. Combine them to replace small loops.

## Example

```go title="slices.go"
s := []int{3, 1, 2}
slices.Sorted(slices.Values(s))   // [1 2 3]
slices.Collect(Countdown(3))      // [3 2 1]
for chunk := range slices.Chunk([]int{1, 2, 3, 4, 5}, 2) {
	fmt.Println(chunk) // [1 2] [3 4] [5]
}
```

| Function | Does |
| --- | --- |
| `All(s)` | Iterator of index, value |
| `Values(s)` | Iterator of values |
| `Backward(s)` | Index, value from the end |
| `Collect(seq)` | Iterator → new slice |
| `AppendSeq(s, seq)` | Append an iterator's values |
| `Sorted(seq)`, `SortedFunc` | Collect and sort |
| `Chunk(s, n)` | Sub-slices of size n |

The older functions (`Sort`, `Contains`, `Index` …) are on [slices Package](/dim-kb/programming/go/arrays-slices/slices-package/).

:::caution[Pitfalls]
- `Chunk` returns views into the original slice, not copies.
:::
