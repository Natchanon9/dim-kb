---
title: "iota"
description: "Auto-incrementing constants for enums and bit sizes"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`iota` is a counter inside a `const (...)` block: 0 on the first line, then +1 per line. Go uses it instead of an `enum` keyword.

## Enums

```go title="iota.go"
type Status int

const (
	Pending Status = iota // 0
	Active                // 1
	Closed                // 2
)
```

Lines without a value repeat the previous expression with the next `iota`.

## Expressions and skipping

```go title="sizes.go"
const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
)
fmt.Println(Pending, Active, Closed, KB, MB) // 0 1 2 1024 1048576
```

:::caution[Pitfalls]
- `iota` restarts at 0 in every new `const` block.
- If 0 means "unset", start real values at `iota + 1` or skip 0 with `_`, so a zero-value field isn't mistaken for a real status.
:::
