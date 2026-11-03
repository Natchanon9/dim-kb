---
title: "recover"
description: "Catch a panic inside a deferred function"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`recover()` stops a panic and returns its value, but only when called directly inside a deferred function.

## Turning a panic into an error

```go title="recover.go"
func safeDivide(a, b int) (q int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

safeDivide(6, 3) // 2 <nil>
safeDivide(1, 0) // 0 recovered: runtime error: integer divide by zero
```

## Real use

HTTP servers recover per request so one bad handler doesn't kill the server (`net/http` does this; frameworks offer recovery middleware). Worker goroutines often recover and report too.

:::caution[Pitfalls]
- `recover` only catches panics in the **same goroutine**; each goroutine you start needs its own.
- Don't use panic/recover as try/catch for normal errors.
:::
