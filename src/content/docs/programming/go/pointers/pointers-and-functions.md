---
title: "Pointers & Functions"
description: "Let a function change the caller's variable, and return pointers safely"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Arguments are copies. Passing a pointer lets the function change the caller's variable.

## Example

```go title="funcs.go"
func reset(n *int) { *n = 0 }

func newInt(v int) *int { return &v } // safe in Go

x := 5
reset(&x)
p := newInt(3)
fmt.Println(x, *p) // 0 3
```

## Returning a pointer to a local

Unlike C, returning `&v` of a local variable is fine. The compiler's escape analysis moves `v` to the heap when it outlives the function.

:::caution[Pitfalls]
- Pointers don't make code faster by default. Small structs passed by value often avoid heap allocations and are faster.
- A function that receives a pointer should document whether it may keep it after returning.
:::
