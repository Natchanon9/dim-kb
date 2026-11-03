---
title: "Interface vs Concrete Types"
description: "When to return a struct, and the typed-nil trap"
sidebar:
  order: 12
status: done
lastReviewed: 2026-11-03
---

Return concrete types so callers get every method and field; accept interfaces so callers can pass anything that fits. The one place this goes wrong is returning a nil pointer as an interface.

## The typed-nil trap

```go title="typednil.go"
type MyErr struct{}

func (*MyErr) Error() string { return "boom" }

func mayFail(fail bool) error {
	var p *MyErr // nil pointer
	if fail {
		p = &MyErr{}
	}
	return p // BUG: non-nil interface holding a nil *MyErr
}

err := mayFail(false)
fmt.Println(err == nil) // false
```

The interface holds `(type=*MyErr, value=nil)`, and an interface is nil only when **both** parts are nil. Fix: `return nil` explicitly on the success path.

## Choosing

| Return a concrete type | Return an interface |
| --- | --- |
| Constructors (`NewServer`) | `error` |
| Callers need fields/methods | Several implementations chosen at runtime |

:::caution[Pitfalls]
- Never return a typed nil pointer through an `error` result.
:::
