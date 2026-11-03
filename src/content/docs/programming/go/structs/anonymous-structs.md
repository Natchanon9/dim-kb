---
title: "Anonymous Structs"
description: "One-off struct types without a name"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

An anonymous struct is declared and used in one place, without a `type` declaration.

## Typical uses

```go title="anon.go"
point := struct{ X, Y int }{1, 2}

tests := []struct { // table-driven test cases
	in   string
	want int
}{
	{"a", 1},
	{"bb", 2},
}
fmt.Println(point, len(tests), tests[1].want) // {1 2} 2 2
```

Also handy for one-off JSON request/response shapes inside a handler, and for grouping config in a variable.

:::caution[Pitfalls]
- If the same shape appears twice, give it a name.
:::
