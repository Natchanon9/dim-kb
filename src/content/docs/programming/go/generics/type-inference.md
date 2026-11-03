---
title: "Type Inference"
description: "When the compiler works out the type arguments for you"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

The compiler infers type arguments from the function arguments, so you rarely write them out.

## Example

```go title="infer.go"
Map([]int{1, 2}, strconv.Itoa) // T=int, U=string inferred → [1 2]

func Zero[T any]() T {
	var z T
	return z
}
Zero[string]() // "" : no arguments, so T must be given
Zero[int]()    // 0
```

## When you must be explicit

- The type parameter appears only in the result (`Zero[T]`, `make`-like constructors).
- Untyped constants would infer a type you don't want: `Larger[float64](1, 2)`.

:::caution[Pitfalls]
- Partial instantiation works left to right: `Map[int]` fixes `T`, then `U` is inferred.
:::
