---
title: "Switch Statement"
description: "switch without break, multiple values per case, tagless switch and fallthrough"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Go's `switch` stops after the first matching case, so no `break` is needed. It is often cleaner than an `if-else` chain.

## Value switch

```go title="switch.go"
func dayKind(day string) string {
	switch day {
	case "sat", "sun": // several values in one case
		return "weekend"
	default:
		return "weekday"
	}
}
```

## Tagless switch and fallthrough

```go title="tagless.go"
x := 15
switch { // same as switch true: first true case wins
case x > 10:
	fmt.Println("big")
	fallthrough // runs the next case body without checking it
case x > 100:
	fmt.Println("huge?")
default:
	fmt.Println("small")
}
// big
// huge?
```

:::caution[Pitfalls]
- `fallthrough` skips the next condition entirely, which is rarely what you want.
- `break` inside a `switch` inside a `for` only leaves the switch. Use a label to leave the loop.
:::
