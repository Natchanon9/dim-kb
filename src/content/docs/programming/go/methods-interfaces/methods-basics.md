---
title: "Methods Basics"
description: "Attach functions to types with a receiver"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A method is a function with a **receiver**, written before the name. You can define methods on any named type declared in your package, not only structs.

## Defining methods

```go title="methods.go"
type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { return r.W * r.H }

type Celsius float64

func (c Celsius) Fahrenheit() float64 { return float64(c)*9/5 + 32 }

fmt.Println(Rect{3, 4}.Area(), Celsius(100).Fahrenheit()) // 12 212
```

The receiver name is short, usually one or two letters of the type (`r`, `c`), never `this` or `self`.

:::caution[Pitfalls]
- You can't add methods to types from other packages (`time.Time`, `int`); define your own named type instead.
:::
