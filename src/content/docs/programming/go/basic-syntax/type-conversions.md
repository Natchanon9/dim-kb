---
title: "Type Conversions"
description: "Explicit conversion with T(x) and strconv for strings"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Go never converts types implicitly. Write `T(x)` to convert between numeric types, and use `strconv` for strings.

## Numbers

```go title="convert.go"
i := 42
f := float64(i)
x := 3.9
fmt.Println(f, int(x)) // 42 3  (int() truncates toward zero)
```

## Strings and numbers

```go title="strconv.go"
s := strconv.Itoa(i)          // "42"
n, err := strconv.Atoi("123") // 123, nil
fmt.Println(s, n, err, string(rune(65))) // 42 123 <nil> A
```

| Need | Use |
| --- | --- |
| int ↔ string | `strconv.Itoa`, `strconv.Atoi` |
| float/bool ↔ string | `strconv.ParseFloat`, `strconv.FormatFloat`, `strconv.ParseBool` |
| Any value → string | `fmt.Sprint(v)` (slower, fine for logs) |

:::caution[Pitfalls]
- `string(65)` gives `"A"`, not `"65"`; `go vet` warns about it. Use `strconv.Itoa`.
- Converting to a smaller integer type silently wraps: `uint8(300)` is a compile error for constants but wraps for variables.
:::
