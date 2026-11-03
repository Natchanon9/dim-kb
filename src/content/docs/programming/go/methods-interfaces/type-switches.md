---
title: "Type Switches"
description: "Handle several concrete types stored in an interface"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

A type switch is a chain of type assertions in one statement. Each case binds the value with that case's type.

## Example

```go title="area.go"
func area(s any) float64 {
	switch v := s.(type) {
	case Rect:
		return v.Area()
	case *Square:
		return v.Area()
	default:
		return 0
	}
}
// area(Rect{2, 3}), area(&Square{3}), area("x") → 6 9 0
```

`case Rect` and `case *Square` are different: a value and a pointer are distinct types.

:::caution[Pitfalls]
- A type switch over your own types is often a missing interface method. If every case calls the same method, call it through the interface instead.
:::
