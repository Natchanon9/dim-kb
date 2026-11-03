---
title: "If-Else"
description: "Conditions without parentheses, and returning early"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`if` takes a boolean condition without parentheses, and the braces are always required.

## Syntax

```go title="grade.go"
func grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	}
	return "C"
}
// grade(85), grade(72), grade(10) → A B C
```

## Return early instead of `else`

Idiomatic Go handles the error or edge case first and returns, so the main path stays unindented:

```go
if err != nil {
	return err
}
// happy path continues here
```

:::caution[Pitfalls]
- The condition must be a `bool`: `if n {}` with an int doesn't compile.
- `else` must be on the same line as the closing `}`.
:::
