---
title: "Basic Types"
description: "Integers, floats, bool, string, byte and rune"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go has a small set of built-in types. Use `int`, `float64`, `string` and `bool` unless you have a reason not to.

## The types

| Kind | Types | Notes |
| --- | --- | --- |
| Integer | `int`, `int8`…`int64`, `uint`…`uint64` | `int` is 64-bit on 64-bit platforms |
| Float | `float32`, `float64` | Default is `float64` |
| Complex | `complex64`, `complex128` | Rarely used |
| Text | `string` | Immutable UTF-8 bytes |
| Aliases | `byte` = `uint8`, `rune` = `int32` | A rune is one Unicode code point |
| Boolean | `bool` | `true` / `false` only, no truthy ints |

```go title="types.go"
var i int = 42
var f float64 = 3.14
var b byte = 'A'
var r rune = 'ก'
var s string = "hi"
fmt.Printf("%T %T %T %T %T\n", i, f, b, r, s)
// int float64 uint8 int32 string
```

:::caution[Pitfalls]
- Mixing types needs an explicit conversion: `int + int64` doesn't compile.
- Floats are inexact: don't store money in `float64`; use integer cents.
:::
