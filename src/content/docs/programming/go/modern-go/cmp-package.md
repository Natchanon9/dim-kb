---
title: "cmp Package"
description: "Compare, Less, Ordered and Or for sorting"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`cmp` (Go 1.21+) provides comparison helpers that pair with `slices.SortFunc`.

## Example

```go title="cmp.go"
cmp.Compare(1, 2)       // -1
cmp.Less("b", "a")      // false
cmp.Or("", "default")   // "default": first non-zero value (1.22+)

slices.SortFunc(staff, func(a, b employee) int {
	return cmp.Or( // by age, then by name
		cmp.Compare(a.Age, b.Age),
		cmp.Compare(a.Name, b.Name),
	)
})
// [{cy 25} {al 30} {bob 30}]
```

| Item | Use |
| --- | --- |
| `cmp.Ordered` | Constraint for `< <= > >=` types |
| `cmp.Compare(a, b)` | -1 / 0 / +1, handles NaN consistently |
| `cmp.Less(a, b)` | Boolean less-than |
| `cmp.Or(vals...)` | First non-zero: multi-key sorts, defaults |

:::caution[Pitfalls]
- `cmp.Or` evaluates every argument; it doesn't short-circuit expensive calls.
:::
