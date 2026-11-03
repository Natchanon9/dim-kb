---
title: "Type Assertions"
description: "Get the concrete value out of an interface with x.(T)"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`x.(T)` extracts the value of type `T` from interface `x`. The two-value form tells you if it worked instead of panicking.

## Two forms

```go title="assert.go"
var v any = "go"
s := v.(string)  // panics if v is not a string
n, ok := v.(int) // safe form: ok reports success
fmt.Println(s, n, ok) // go 0 false
```

## Asserting to an interface

`T` can be another interface: `if s, ok := v.(fmt.Stringer); ok { ... }` checks whether the value has a `String` method. Standard libraries use this to detect optional capabilities (e.g. `io.WriterTo`).

:::caution[Pitfalls]
- Use the one-value form only when a wrong type is a programming bug; otherwise use `, ok`.
- For errors, use `errors.As` instead of asserting, so wrapped errors are found.
:::
