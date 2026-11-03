---
title: "Method Sets"
description: "Which methods a value or pointer has, and how that affects interfaces"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

The method set decides which interfaces a type satisfies.

| Type | Method set |
| --- | --- |
| `T` | Methods with value receivers |
| `*T` | Methods with value **and** pointer receivers |

## Why it matters

```go title="sets.go"
type Shape interface{ Area() float64 }

type Square struct{ S float64 }

func (s *Square) Area() float64 { return s.S * s.S } // pointer receiver

var sh Shape = &Square{2} // Square{2} would not compile
fmt.Println(sh.Area())    // 4
```

Calling `sq.Area()` on an addressable `Square` variable still works, because Go takes the address for you. Storing a `Square` value in an interface doesn't, because the interface holds a copy that has no address.

:::caution[Pitfalls]
- The error "Square does not implement Shape (method Area has pointer receiver)" means: pass `&Square{}`.
:::
