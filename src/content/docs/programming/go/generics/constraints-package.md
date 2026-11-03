---
title: "constraints Package"
description: "comparable, cmp.Ordered and the old x/exp/constraints"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

You need very few predefined constraints. Since Go 1.21, `cmp.Ordered` in the standard library replaced most uses of `golang.org/x/exp/constraints`.

## Example

```go title="keys.go"
func Index[T comparable](s []T, v T) int { // comparable: == and !=
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K { // cmp.Ordered: < >
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
// Index([]string{"a","b"}, "b") → 1 ; SortedKeys → [a b]
```

| Need | Use |
| --- | --- |
| Equality, map keys | `comparable` (built in) |
| Ordering | `cmp.Ordered` |
| Integers only, floats only | Define your own with `~int \| ~int64 ...` |

:::caution[Pitfalls]
- `x/exp/constraints` is experimental; prefer `cmp` in new code.
:::
