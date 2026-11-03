---
title: "Variables & Declarations"
description: "Declaring variables with var, type inference and grouped declarations"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`var` declares a variable with a name, a type, or both. Go is statically typed: a variable's type never changes.

## Three forms of `var`

```go title="vars.go"
var name string = "gopher"
var age = 3 // type inferred as int
var (
	host = "localhost"
	port = 8080
)
fmt.Println(name, age, host, port) // gopher 3 localhost 8080
```

- **Type + value:** use when the inferred type would be wrong (`var ratio float64 = 1`).
- **Value only:** the type comes from the value.
- **Grouped `var (...)`:** common for package-level variables.

## Scope

A variable is visible inside the `{}` block where it is declared. Package-level variables are visible to the whole package; they can't use `:=`.

:::caution[Pitfalls]
- An unused local variable is a compile error.
- Declaring the same name in an inner block **shadows** the outer one; changes inside don't affect the outer variable.
:::
