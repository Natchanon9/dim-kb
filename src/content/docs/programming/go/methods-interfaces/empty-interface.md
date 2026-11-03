---
title: "Empty Interface"
description: "any / interface{} holds a value of any type"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`any` (alias of `interface{}`, Go 1.18+) has no methods, so every type satisfies it. It can hold any value.

## Example

```go title="any.go"
var v any = 42
v = "hi"
v = []int{1}
fmt.Printf("%T\n", v) // []int

values := []any{1, "a", true, nil}
fmt.Println(values...) // 1 a true <nil>
```

## When to use it

- Truly unknown data: decoding arbitrary JSON (`map[string]any`), `fmt.Println`, logging fields.
- To get the value back, use a type assertion or type switch.

Prefer generics or a real interface when the set of types is known.

:::caution[Pitfalls]
- `any` throws away compile-time type checks; mistakes show up as runtime panics.
- JSON numbers decode into `any` as `float64`, not `int`.
:::
