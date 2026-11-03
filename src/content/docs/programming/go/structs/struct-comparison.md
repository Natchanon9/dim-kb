---
title: "Struct Comparison"
description: "== on structs, map keys, and comparing structs with slices"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Structs are comparable with `==` when every field is comparable. Comparable structs can also be map keys.

## Example

```go title="compare.go"
type Point struct{ X, Y int }

Point{1, 2} == Point{1, 2}           // true
seen := map[Point]bool{{1, 2}: true} // struct as map key

type Tagged struct {
	Name string
	Tags []string // slice field: not comparable
}
a := Tagged{"x", []string{"go"}}
b := Tagged{"x", []string{"go"}}
reflect.DeepEqual(a, b) // true; a == b would not compile
```

## Comparing non-comparable structs

- Tests: `reflect.DeepEqual` or `github.com/google/go-cmp`.
- Production: write an `Equal` method; it is faster and states what equality means.

:::caution[Pitfalls]
- Comparing pointers (`&a == &b`) checks identity, not content.
- Adding a slice or map field later breaks every `==` on that type at compile time.
:::
