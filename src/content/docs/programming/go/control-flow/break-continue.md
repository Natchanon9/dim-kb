---
title: "Break & Continue"
description: "Leave a loop early or skip to the next iteration"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`break` leaves the innermost `for`, `switch` or `select`. `continue` jumps to the next iteration of the innermost loop.

## Example

```go title="skip.go"
for i := range 10 {
	if i%2 == 0 {
		continue // skip even numbers
	}
	if i > 7 {
		break // stop the loop
	}
	fmt.Print(i, " ")
}
// 1 3 5 7
```

:::caution[Pitfalls]
- Inside `for { select { ... } }` or `for { switch { ... } }`, a plain `break` only exits the `select`/`switch`. The loop keeps running. Use a labeled break or `return`.
:::
