---
title: "Type Constraints"
description: "Interfaces that list which types or operations are allowed"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

A constraint is an interface. Besides methods, it can list allowed types with `|`, which unlocks operators like `+` and `<` inside the function.

## Example

```go title="sum.go"
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n // allowed: every type in Number supports +
	}
	return total
}

type Score int
Sum([]int{1, 2})       // 3
Sum([]float64{1.5, 2}) // 3.5
Sum([]Score{3, 4})     // 7, thanks to ~int
```

## Built-in constraints

| Constraint | Allows |
| --- | --- |
| `any` | Every type; only assignment and passing around |
| `comparable` | `==`, `!=`; usable as map key |
| `cmp.Ordered` | `< <= > >=`, numbers and strings |

:::caution[Pitfalls]
- Interfaces with a type list (`int | float64`) can only be used as constraints, not as ordinary variable types.
:::
