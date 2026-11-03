---
title: "Anonymous Functions"
description: "Function literals without a name"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A function literal has no name. Write it inline where a function value is needed, or call it immediately.

## Common uses

```go title="anon.go"
result := func(a, b int) int { return a * b }(3, 4) // defined and called: 12

names := []string{"bob", "al", "christine"}
sort.Slice(names, func(i, j int) bool { return len(names[i]) < len(names[j]) })
// [al bob christine]
```

You will write them for sort comparators, `defer func() { ... }()`, `go func() { ... }()` and HTTP handlers.

:::caution[Pitfalls]
- An anonymous function that uses outer variables is a closure: it sees their **current** value when it runs, not the value when it was created.
- Long literals hurt readability; give them a name once they grow past a few lines.
:::
