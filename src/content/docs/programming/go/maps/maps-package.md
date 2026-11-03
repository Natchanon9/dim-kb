---
title: "maps Package"
description: "Clone, compare, copy and filter maps (Go 1.21+)"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

The generic `maps` package (Go 1.21+) covers the map loops you'd otherwise write by hand.

## Common functions

```go title="pkg.go"
a := map[string]int{"x": 1}
b := maps.Clone(a)
b["y"] = 2
fmt.Println(maps.Equal(a, b), len(a), len(b)) // false 1 2

maps.Copy(a, b) // copy b's entries into a
maps.DeleteFunc(a, func(k string, v int) bool { return v > 1 })
fmt.Println(a) // map[x:1]
```

| Function | Does |
| --- | --- |
| `Clone` | Shallow copy |
| `Equal`, `EqualFunc` | Compare contents |
| `Copy(dst, src)` | Insert/overwrite all of src into dst |
| `DeleteFunc` | Delete entries matching a predicate |
| `Keys`, `Values` (1.23+) | Iterators over keys / values |

:::caution[Pitfalls]
- `Clone` is shallow: slice or pointer values are shared between the two maps.
:::
