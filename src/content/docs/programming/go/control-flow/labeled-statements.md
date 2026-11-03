---
title: "Labeled Statements"
description: "break and continue an outer loop with a label"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

A label names a loop so `break` or `continue` can target it from inside nested loops or a `select`.

## Example

```go title="labels.go"
outer:
	for i := range 3 {
		for j := range 3 {
			if j == 2 {
				continue outer // next i
			}
			if i == 2 {
				break outer // leave both loops
			}
			fmt.Printf("%d%d ", i, j)
		}
	}
// 00 01 10 11
```

The most common real use is leaving a `for { select { ... } }` loop when a channel closes.

:::caution[Pitfalls]
- An unused label is a compile error.
- If you need labels often, extracting the inner loop into a function with `return` is usually clearer.
:::
