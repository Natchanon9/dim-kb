---
title: "Constructor Functions"
description: "NewXxx functions that validate input and set defaults"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Go has no constructors. By convention a function named `NewT` builds a `T` (often `*T`), validates input and fills in defaults.

## Example

```go title="server.go"
type Server struct {
	addr    string
	timeout time.Duration
}

func NewServer(addr string) (*Server, error) {
	if addr == "" {
		return nil, errors.New("addr required")
	}
	return &Server{addr: addr, timeout: 30 * time.Second}, nil
}
```

## When you need one

- Required fields or validation.
- Defaults other than zero values.
- Internal state that must be initialised (maps, channels).

If the zero value works, skip the constructor.

:::caution[Pitfalls]
- Many optional settings → use an options struct or functional options, not ten parameters.
- A constructor that can fail should return an `error`, not panic.
:::
