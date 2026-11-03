---
title: "Slice Tricks"
description: "Delete, insert, reverse and filter in place"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Common edits are one-liners with the `slices` package (Go 1.21+). Filtering in place reuses the backing array to avoid an allocation.

## Delete, insert, reverse

```go title="tricks.go"
s := []int{1, 2, 3, 4, 5}
s = slices.Delete(s, 1, 2) // remove index 1 → [1 3 4 5]
s = slices.Insert(s, 1, 2) // put 2 back     → [1 2 3 4 5]
slices.Reverse(s)          //                → [5 4 3 2 1]
```

## Filter in place

```go title="filter.go"
evens := s[:0] // same backing array, length 0
for _, v := range s {
	if v%2 == 0 {
		evens = append(evens, v)
	}
}
// evens → [4 2]
```

Or use `slices.DeleteFunc(s, func(v int) bool { return v%2 != 0 })`.

:::caution[Pitfalls]
- Filtering in place overwrites the original slice. Don't use `s` afterwards, only the result.
- `Delete` and `Insert` return a new header; always assign the result.
:::
