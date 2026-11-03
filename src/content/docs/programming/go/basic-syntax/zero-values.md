---
title: "Zero Values"
description: "Every variable starts with a usable default"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A declared variable without a value gets its type's **zero value**. There is no "uninitialised" state in Go.

## Zero value per type

| Type | Zero value |
| --- | --- |
| Numbers | `0` |
| `bool` | `false` |
| `string` | `""` |
| Pointer, slice, map, channel, func, interface | `nil` |
| Struct | Each field set to its zero value |

```go title="zero.go"
var i int
var s string
var ok bool
var p *int
var list []int
var m map[string]int
fmt.Printf("%d %q %t %v %t %t\n", i, s, ok, p, list == nil, m == nil)
// 0 "" false <nil> true true
```

## Make the zero value useful

Good Go types work without a constructor: `sync.Mutex`, `bytes.Buffer` and a nil slice (you can `append` to it) are all ready at their zero value.

:::caution[Pitfalls]
- Writing to a nil map panics. Create it with `make(map[K]V)` first.
- Calling a method through a nil pointer panics if the method reads fields.
:::
