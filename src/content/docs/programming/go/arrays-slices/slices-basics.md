---
title: "Slices Basics"
description: "Dynamic views over arrays: literal, index, slice expression, len"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A slice is a growable view over an array. It is the list type you use everywhere in Go.

## Basics

```go title="slice.go"
s := []int{1, 2, 3} // no length in the brackets → slice
s = append(s, 4)
fmt.Println(s, len(s), s[1:3]) // [1 2 3 4] 4 [2 3]
```

- `s[i]` reads/writes one element.
- `s[low:high]` makes a new slice from `low` up to, not including, `high`. Omit either end: `s[:2]`, `s[1:]`.
- `len(s)` is the element count; `cap(s)` is room before a reallocation.

## Slices vs arrays

| | Array `[3]int` | Slice `[]int` |
| --- | --- | --- |
| Length | Fixed, part of the type | Grows with `append` |
| Assignment | Copies elements | Copies the header, shares data |
| Comparable with `==` | Yes | Only to `nil` |

:::caution[Pitfalls]
- Always use the result of `append`: `s = append(s, x)`.
- Sub-slices share memory with the original; writes through one are visible in the other.
:::
