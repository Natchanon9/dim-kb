---
title: "Multiple Return Values"
description: "Return a result and an error together"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A Go function can return several values. The standard shape is `(result, error)`, with the error last.

## The `(value, error)` pattern

```go title="divide.go"
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

q, err := divide(10, 4) // 2.5 <nil>
_, err = divide(1, 0)   // divide by zero
```

Check `err` first; when it is non-nil, treat the other results as meaningless.

:::caution[Pitfalls]
- You must receive all results or none: `x := divide(1, 2)` doesn't compile. Use `_` to drop one.
- Returning more than 2–3 values is a sign you want a struct.
:::
