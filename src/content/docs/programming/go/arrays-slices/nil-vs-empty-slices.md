---
title: "Nil vs Empty Slices"
description: "Both have length 0, but they differ in == nil and in JSON"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A nil slice and an empty slice both have length 0 and both work with `len`, `range` and `append`. They differ in two places.

## The difference

```go title="nil.go"
var a []int  // nil
b := []int{} // empty, not nil
fmt.Println(a == nil, b == nil, len(a), len(b)) // true false 0 0

ja, _ := json.Marshal(a) // null
jb, _ := json.Marshal(b) // []
```

## Which to use

- Inside code, prefer `var s []T` and check `len(s) == 0`, never `s == nil`.
- In API responses where clients expect `[]`, initialise with `[]T{}` (or `make([]T, 0)`).

:::caution[Pitfalls]
- A JSON field that is sometimes `null` and sometimes `[]` breaks strict clients. Pick one on purpose.
:::
