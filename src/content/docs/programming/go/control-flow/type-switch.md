---
title: "Type Switch"
description: "Branch on the dynamic type of an interface value"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A type switch checks which concrete type is stored in an interface value, and gives you a variable of that type in each case.

## Syntax

```go title="describe.go"
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int %d", x+1) // x is an int here
	case string:
		return "string " + x
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("other %T", x)
	}
}
// describe(1) → int 2, describe(2.5) → other float64
```

In a case listing several types (`case int, int64:`), `x` keeps the interface type.

:::caution[Pitfalls]
- It only works on interface values (`any`, `error`, `io.Reader` …), not on concrete types.
- Prefer a method on an interface over a big type switch; adding a type then doesn't mean editing every switch.
:::
