---
title: "Variadic Functions"
description: "Accept any number of arguments with ...T"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A final parameter `...T` accepts zero or more arguments. Inside the function it is a `[]T`.

## Declaring and spreading

```go title="sum.go"
func sum(nums ...int) int { // nums is a []int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

sum()        // 0
sum(1, 2, 3) // 6
xs := []int{4, 5}
sum(xs...)   // 9, spread a slice
```

`fmt.Println(a ...any)` and `append(s, items...)` are the variadic functions you use every day.

:::caution[Pitfalls]
- Only the last parameter can be variadic.
- `f(xs...)` passes the same slice, not a copy; the function can modify your elements.
:::
