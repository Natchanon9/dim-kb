---
title: "Creating Slices"
description: "Literal, make with length and capacity, nil slice, slicing arrays"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Pick the creation form by what you know up front: the values, the size, or nothing yet.

## Forms

```go title="create.go"
a := []int{1, 2}        // literal
b := make([]int, 3)     // len 3, all zero
c := make([]int, 0, 10) // len 0, room for 10
var d []int             // nil slice
fmt.Println(a, b, len(c), cap(c), d == nil)
// [1 2] [0 0 0] 0 10 true
```

| You know | Use |
| --- | --- |
| The values | Literal `[]T{...}` |
| The exact length | `make([]T, n)`, then assign `s[i]` |
| Roughly how many you'll append | `make([]T, 0, n)`, then `append` |
| Nothing yet | `var s []T` |

:::caution[Pitfalls]
- `make([]T, n)` then `append` adds **after** the n zeros. Use `make([]T, 0, n)` if you plan to append.
:::
