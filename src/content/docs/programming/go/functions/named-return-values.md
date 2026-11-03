---
title: "Named Return Values"
description: "Name the results, naked return, and editing results in defer"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Results can have names. They start at their zero value and a bare `return` returns their current values.

## Naked return

```go title="split.go"
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // returns x, y
}
// split(17) → 7 10
```

## Where names really help

A deferred function can read and change named results, which is the standard way to add context to every returned error:

```go title="parse.go"
func parse(s string) (n int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse %q: %w", s, err)
		}
	}()
	return strconv.Atoi(s)
}
// parse("x") → parse "x": strconv.Atoi: parsing "x": invalid syntax
```

Names also document what each result means: `(width, height int)`.

:::caution[Pitfalls]
- Naked returns in long functions hide what is returned; use them only in short ones.
- `x, err := ...` inside an inner block shadows the named result; the outer one is what gets returned.
:::
