---
title: "Arrays"
description: "Fixed-length arrays and why you rarely use them directly"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

An array has a fixed length that is part of its type: `[3]int` and `[4]int` are different types. Arrays are values, so assigning one copies every element.

## Declaring and copying

```go title="arrays.go"
var a [3]int
b := [...]string{"x", "y"} // length from the literal
a[0] = 1
c := a // copies all elements
c[0] = 99
fmt.Println(a, len(b), c, a == [3]int{1, 0, 0})
// [1 0 0] 2 [99 0 0] true
```

## When to use arrays

Mostly you use slices. Arrays fit fixed-size data: a SHA-256 sum (`[32]byte`), RGB colours, or as comparable map keys (`map[[2]int]bool`).

:::caution[Pitfalls]
- Passing a large array to a function copies it. Pass a slice or a pointer instead.
- Index out of range is a compile error for constants and a panic at runtime otherwise.
:::
