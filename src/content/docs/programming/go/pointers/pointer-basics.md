---
title: "Pointer Basics"
description: "& takes an address, * reads or writes through it"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A pointer holds the memory address of a value. `&x` takes the address; `*p` reads or writes the value it points to.

## Example

```go title="pointer.go"
x := 10
p := &x // p is a *int holding x's address
*p = 20 // write through the pointer
fmt.Println(x, *p, p != nil) // 20 20 true
```

| Syntax | Meaning |
| --- | --- |
| `*T` | Type "pointer to T" |
| `&x` | Address of x |
| `*p` | Value p points to (dereference) |
| `nil` | Zero value of any pointer |

Go has no pointer arithmetic (outside `unsafe`), and the garbage collector frees values once nothing points to them.

:::caution[Pitfalls]
- Dereferencing a nil pointer panics: "invalid memory address or nil pointer dereference".
:::
