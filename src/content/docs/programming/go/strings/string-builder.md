---
title: "String Builder"
description: "Build strings efficiently with strings.Builder"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Concatenating with `+` in a loop copies the whole string every time. `strings.Builder` appends into a growing buffer instead.

## Example

```go title="builder.go"
var sb strings.Builder // zero value is ready to use
for i := range 3 {
	fmt.Fprintf(&sb, "%d,", i) // Builder is an io.Writer
}
sb.WriteString("done")
fmt.Println(sb.String(), sb.Len()) // 0,1,2,done 10
```

Methods: `WriteString`, `WriteByte`, `WriteRune`, `Write`, `Grow(n)` to pre-size, `Reset`.

## When to bother

- A few `+` outside loops: fine.
- Joining a slice: `strings.Join`.
- Building in a loop: `strings.Builder`.

:::caution[Pitfalls]
- Don't copy a `Builder` after writing to it; pass `*strings.Builder`.
:::
