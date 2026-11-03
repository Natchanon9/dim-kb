---
title: "Dynamic Invocation"
description: "Call methods and functions by name at runtime"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`Value.MethodByName` finds a method at runtime and `Value.Call` invokes it with `[]reflect.Value` arguments.

## Example

```go title="call.go"
type Calc struct{}

func (Calc) Add(a, b int) int { return a + b }

m := reflect.ValueOf(Calc{}).MethodByName("Add")
out := m.Call([]reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)})
fmt.Println(out[0].Int()) // 5
```

Used by RPC frameworks, plugin systems and command routers that map names to methods.

## Usually better

A `map[string]func(...)` gives the same dispatch with compile-time checks and no reflection.

:::caution[Pitfalls]
- `MethodByName` returns an invalid `Value` for a missing or unexported method; check `m.IsValid()` before `Call`.
- Wrong argument count or types panic.
:::
