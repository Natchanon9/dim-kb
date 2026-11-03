---
title: "Struct Literals"
description: "Create structs with named, positional and pointer literals"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A struct literal creates a value with some or all fields set. Fields you omit get their zero value.

## Three forms

```go title="literals.go"
a := User{Name: "Bob", Age: 20} // named fields: preferred
b := User{"Cat", 30}            // positional: every field, in order
p := &User{Name: "Dan"}         // pointer to a new struct
p.Age = 40                      // same as (*p).Age
fmt.Println(a, b, *p) // {Bob 20} {Cat 30} {Dan 40}
```

Use named fields: they survive adding or reordering fields. `go vet` warns about positional literals of types from other packages.

:::caution[Pitfalls]
- A trailing comma is required when the closing `}` is on its own line.
:::
