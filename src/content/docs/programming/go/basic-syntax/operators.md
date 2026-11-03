---
title: "Operators"
description: "Arithmetic, comparison, logical and bitwise operators"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go's operators look like C's, with fewer surprises: no ternary, and `++`/`--` are statements.

## Arithmetic and bitwise

```go title="ops.go"
fmt.Println(7/2, 7%2, 7.0/2)           // 3 1 3.5
fmt.Println(1<<3, 6&3, 6|3, 6^3, 6&^3) // 8 2 7 5 4
x := 5
x += 2
x++ // a statement, not an expression
fmt.Println(x) // 8
```

| Group | Operators |
| --- | --- |
| Arithmetic | `+ - * / %` (integer `/` truncates) |
| Comparison | `== != < <= > >=` |
| Logical | `&& \|\| !` (short-circuit) |
| Bitwise | `& \| ^ &^ << >>` (`&^` = AND NOT, clears bits) |
| Assignment | `= += -= *= /= %= <<= >>= &= \|= ^=` |

:::caution[Pitfalls]
- No `?:` ternary; write an `if`.
- `y := x++` doesn't compile.
- Integer division by a zero variable panics at runtime.
:::
