---
title: "Struct Basics"
description: "Group related fields into a named type"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A struct groups named fields into one type. It is Go's main way to model data; there are no classes.

## Defining and using

```go title="user.go"
type User struct {
	Name string
	Age  int
}

var u User // all fields at their zero value
u.Name = "Ann"
fmt.Println(u, u.Age) // {Ann 0} 0
```

Fields are accessed with `.`, both on values and on pointers (`p.Name` works when `p` is `*User`).

## Value semantics

Assigning or passing a struct copies all its fields. Use a pointer (`*User`) when a function must modify it or the struct is large.

:::caution[Pitfalls]
- Changing a copy doesn't change the original: `for _, u := range users { u.Age++ }` does nothing to `users`.
:::
