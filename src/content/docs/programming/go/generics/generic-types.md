---
title: "Generic Types"
description: "Structs and other types with type parameters"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Types can take type parameters too. Methods then use the type's parameters but can't add new ones.

## Example: Stack

```go title="stack.go"
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

var s Stack[string] // instantiate with a concrete type
s.Push("a")
s.Push("b")
s.Pop() // b true
```

Good fits: containers (stack, queue, set, LRU cache), typed wrappers (`Result[T]`, `Page[T]` for paginated API responses).

:::caution[Pitfalls]
- `var zero T` is the way to return "nothing" of an unknown type.
:::
