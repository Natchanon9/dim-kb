---
title: "Generics Best Practices"
description: "When generics help and when an interface is better"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

Use generics when the same code would otherwise be copied for several types. If the function only calls methods, an ordinary interface is simpler.

## Interface vs generic

```go title="choose.go"
// An interface is enough: the function only calls methods.
func countBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}

// Generics fit: the algorithm needs the element type itself.
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}
```

## Rules of thumb

- **Write the concrete version first;** generalise when a second type needs it.
- **Good fits:** containers, slice/map helpers, typed API wrappers.
- **Poor fits:** business logic, anything already expressed as an interface, "generic repository" layers over ORMs.
- **Keep constraints minimal:** ask for `any` or `comparable` unless you need operators.

:::caution[Pitfalls]
- Type parameters make signatures harder to read; don't add them for one caller.
:::
