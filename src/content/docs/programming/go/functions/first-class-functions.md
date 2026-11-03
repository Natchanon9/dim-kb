---
title: "First-Class Functions"
description: "Functions as values: variables, parameters, map values and types"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Functions are values. You can store them in variables, pass them as arguments, return them, and keep them in maps.

## Passing a function

```go title="apply.go"
func apply(nums []int, f func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

square := func(n int) int { return n * n }
apply([]int{1, 2, 3}, square) // [1 4 9]
```

## Dispatch table

```go title="ops.go"
ops := map[string]func(a, b int) int{
	"add": func(a, b int) int { return a + b },
	"mul": func(a, b int) int { return a * b },
}
ops["mul"](3, 4) // 12
```

Give common signatures a name: `type HandlerFunc func(http.ResponseWriter, *http.Request)`.

:::caution[Pitfalls]
- The zero value of a func is `nil`; calling it panics.
- Funcs can only be compared to `nil`, so they can't be map keys.
:::
