---
title: "Sets with Maps"
description: "Use map[T]struct{} as a set"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Go has no set type. A map with an empty-struct value works as one, and `struct{}` takes zero bytes.

## Example

```go title="set.go"
seen := map[string]struct{}{}
for _, w := range []string{"go", "rust", "go"} {
	seen[w] = struct{}{}
}
_, ok := seen["go"]
fmt.Println(len(seen), ok) // 2 true
```

## `struct{}` vs `bool`

| Value type | Pros |
| --- | --- |
| `map[T]struct{}` | No memory per value; clearly a set |
| `map[T]bool` | Simpler check: `if seen[w] { ... }` |

Both are fine; teams usually pick one style.

:::caution[Pitfalls]
- Union, intersection and difference are loops you write yourself (or a small generic helper).
:::
