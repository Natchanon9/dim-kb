---
title: "Dependency Injection"
description: "Pass dependencies in through constructors"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

In Go, DI means: a type receives what it needs as constructor parameters (usually interfaces), and `main` wires everything together. No framework is needed.

## Example

```go title="greeter.go"
type Clock interface{ Now() time.Time }

type Greeter struct{ clock Clock }

func NewGreeter(c Clock) *Greeter { return &Greeter{clock: c} }

func (g *Greeter) Greet() string {
	if g.clock.Now().Hour() < 12 {
		return "good morning"
	}
	return "good afternoon"
}

// in tests: a fixed clock → "good morning"
```

## Wiring in main

```go
db := mustOpenDB(cfg.DatabaseURL)
repo := order.NewRepository(db)
svc := order.NewService(repo, clock.Real{})
handler := order.NewHandler(svc)
```

Code generators (Wire) or containers (Fx) help only in very large apps.

:::caution[Pitfalls]
- Avoid package-level globals for dependencies (`var DB *sql.DB`); they hide coupling and make tests share state.
:::
