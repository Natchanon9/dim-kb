---
title: "Reflection Basics"
description: "Inspect types and values at runtime with reflect"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Reflection lets a program look at a value's type and contents at runtime. It is how `encoding/json`, `fmt` and ORMs handle types they have never seen.

## Example

```go title="reflect.go"
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println(v.Type(), v.Kind(), v.Float()) // float64 float64 3.4
```

## The two entry points

| Function | Gives |
| --- | --- |
| `reflect.TypeOf(x)` | `reflect.Type`: name, kind, fields, methods |
| `reflect.ValueOf(x)` | `reflect.Value`: the data, with getters and setters |

## Should you use it?

Rarely in application code. Reach for interfaces or generics first. Reflection fits libraries: encoders, validators, dependency injection, test helpers.

:::caution[Pitfalls]
- Calling the wrong getter (`v.Int()` on a string) panics at runtime; check `v.Kind()` first.
:::
