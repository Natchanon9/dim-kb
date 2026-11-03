---
title: "reflect.Type & reflect.Value"
description: "Type vs Kind, and setting values through a pointer"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`Type` describes **what** something is; `Value` holds **the data**. `Kind` is the underlying category (int, struct, slice …), while `Type` is the declared name.

## Example

```go title="typevalue.go"
type Celsius float64

t := reflect.TypeOf(Celsius(1))
fmt.Println(t.Name(), t.Kind()) // Celsius float64

x := 1
v := reflect.ValueOf(&x).Elem() // through a pointer → settable
fmt.Println(v.CanSet()) // true
v.SetInt(42)
fmt.Println(x) // 42
```

## Useful methods

| On `Type` | On `Value` |
| --- | --- |
| `Name`, `Kind`, `String` | `Kind`, `Type`, `Interface()` |
| `NumField`, `Field(i)` | `Field(i)`, `FieldByName` |
| `Elem` (pointer/slice target) | `Elem`, `Len`, `Index(i)` |
| `NumMethod`, `Method(i)` | `Set`, `SetInt`, `SetString` |

:::caution[Pitfalls]
- `reflect.ValueOf(x)` holds a copy; to modify, pass `&x` and call `.Elem()`.
- Unexported fields can be read but not set.
:::
