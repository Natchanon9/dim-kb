---
title: "append & copy"
description: "Grow a slice with append and duplicate data with copy"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`append` adds elements and may allocate a new backing array. `copy` moves elements between existing slices.

## append

```go title="append.go"
s := make([]int, 0, 2)
s = append(s, 1, 2)
before := cap(s)
s = append(s, 3) // over capacity: new, bigger backing array
fmt.Println(s, before, cap(s) > before) // [1 2 3] 2 true
```

Join slices with `append(a, b...)`.

## copy

```go title="copy.go"
dst := make([]int, 2)
n := copy(dst, []int{7, 8, 9}) // copies min(len(dst), len(src))
fmt.Println(dst, n)            // [7 8] 2
```

To duplicate a whole slice, `slices.Clone(s)` is shortest.

:::caution[Pitfalls]
- `copy` into a zero-length slice copies nothing: size `dst` with `make([]T, len(src))` first.
- Appending inside a hot loop without pre-sizing causes repeated allocations; use `make([]T, 0, n)`.
:::
