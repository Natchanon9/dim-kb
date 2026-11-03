---
title: "slices Package"
description: "Sort, search and compare slices with the standard library (Go 1.21+)"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

The generic `slices` package (Go 1.21+) replaces most hand-written loops over slices.

## Common functions

```go title="pkg.go"
s := []int{3, 1, 2}
slices.Sort(s)
i, found := slices.BinarySearch(s, 2)
fmt.Println(s, slices.Contains(s, 3), slices.Index(s, 3), i, found, slices.Max(s))
// [1 2 3] true 2 1 true 3
```

| Need | Function |
| --- | --- |
| Sort | `Sort`, `SortFunc`, `SortStableFunc` |
| Find | `Contains`, `Index`, `IndexFunc`, `BinarySearch` |
| Compare | `Equal`, `Compare` |
| Edit | `Insert`, `Delete`, `DeleteFunc`, `Compact`, `Reverse` |
| Copy | `Clone` |
| Min / max | `Min`, `Max` |

`SortFunc` takes a comparator returning negative/zero/positive, typically `cmp.Compare(a.Age, b.Age)`.

:::caution[Pitfalls]
- `BinarySearch` only works on a sorted slice.
- `Min`/`Max` panic on an empty slice.
:::
