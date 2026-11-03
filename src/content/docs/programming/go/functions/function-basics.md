---
title: "Function Basics"
description: "Declaring functions, parameters and pass-by-value"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A function is declared with `func`, its parameters, and its result types. Every argument is passed **by value**: the function gets a copy.

## Declaring and calling

```go title="add.go"
func add(a, b int) int { return a + b } // a, b share the type

func double(n int) { n *= 2 } // n is a copy

x := 5
double(x)
fmt.Println(add(2, 3), x) // 5 5  (x unchanged)
```

## Pass by value

To let a function change the caller's variable, pass a pointer (`*int`). Slices, maps and channels hold a pointer inside, so a function can change their **elements** without a pointer.

:::caution[Pitfalls]
- No default parameter values and no overloading. Use separate names (`NewServer`, `NewServerWithTLS`) or an options struct.
- Unused parameters are allowed; unused local variables are not.
:::
