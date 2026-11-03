---
title: "Constants"
description: "const values, typed vs untyped constants"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

A `const` is fixed at compile time. Only numbers, strings, booleans (and runes) can be constants.

## Declaring constants

```go title="const.go"
const Pi = 3.14159 // untyped constant

const timeout = 5 * time.Second
var f32 float32 = Pi // untyped Pi fits any float type
fmt.Println(Pi, f32, timeout) // 3.14159 3.14159 5s
```

## Untyped vs typed

- **Untyped** (`const Pi = 3.14`): has a default type but adapts to where it is used, with high precision.
- **Typed** (`const Pi float64 = 3.14`): behaves exactly like a `float64` value.

Prefer untyped constants unless you are defining an enum-like type.

:::caution[Pitfalls]
- Slices, maps and structs can't be `const`; use a package-level `var`.
- A constant that overflows its target type is a compile error, not a runtime wrap-around.
:::
