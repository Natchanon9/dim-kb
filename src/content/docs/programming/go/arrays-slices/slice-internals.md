---
title: "Slice Internals"
description: "The slice header: pointer, length and capacity"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

A slice value is a small header with three fields: a **pointer** into a backing array, a **length**, and a **capacity**. Copying a slice copies the header, not the data.

## Sharing the backing array

```go title="internals.go"
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3] // pointer to arr[1], len 2, cap 4
fmt.Println(s, len(s), cap(s)) // [2 3] 2 4
s[0] = 20 // writes into arr
fmt.Println(arr) // [1 20 3 4 5]
```

Capacity runs from the slice's start to the end of the backing array, which is why `cap(s)` is 4.

## Growth

When `append` exceeds capacity, Go allocates a bigger array (roughly double for small slices, smaller growth factors for large ones), copies the elements, and returns a header pointing to the new array.

:::caution[Pitfalls]
- After a reallocation, the old and new slices no longer share data; before it, they do. Code that depends on either is fragile.
- A tiny sub-slice of a huge slice keeps the whole backing array alive. Copy it out with `slices.Clone` if you keep it long-term.
:::
