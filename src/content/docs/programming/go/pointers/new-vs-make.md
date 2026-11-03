---
title: "new vs make"
description: "new allocates a zeroed value; make builds slices, maps and channels"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Both allocate, but for different things.

| | `new(T)` | `make(T, args)` |
| --- | --- | --- |
| Works on | Any type | Only slice, map, chan |
| Returns | `*T` to a zero value | A ready-to-use `T` (not a pointer) |
| Typical use | Rare; `&T{}` is more common | Always for maps and channels |

## Example

```go title="newmake.go"
p := new(int) // *int pointing to a zero int
*p = 7
s := make([]int, 2)
m := make(map[string]int)
ch := make(chan int, 1)
m["a"] = 1
ch <- 9
fmt.Println(*p, s, m, <-ch) // 7 [0 0] map[a:1] 9

ps := new([]int) // pointer to a nil slice: rarely what you want
fmt.Println(*ps == nil) // true
```

:::caution[Pitfalls]
- `new(map[string]int)` gives a pointer to a nil map; writing to it panics. Use `make`.
:::
