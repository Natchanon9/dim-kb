---
title: "strconv Package"
description: "Convert between strings and numbers or booleans"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`strconv` converts strings to and from numbers and booleans, returning an error for bad input.

## Parse and format

```go title="strconv.go"
n, _ := strconv.Atoi("42")
f, _ := strconv.ParseFloat("3.5", 64)
b, _ := strconv.ParseBool("true")
fmt.Println(n+1, f*2, b, strconv.Itoa(7)+"!", strconv.FormatFloat(1.5, 'f', 2, 64), strconv.Quote("hi"))
// 43 7 true 7! 1.50 "hi"

_, err := strconv.Atoi("12a")
// strconv.Atoi: parsing "12a": invalid syntax
```

| From string | To string |
| --- | --- |
| `Atoi`, `ParseInt(s, base, bits)` | `Itoa`, `FormatInt(n, base)` |
| `ParseFloat(s, 64)` | `FormatFloat(f, 'f', prec, 64)` |
| `ParseBool` | `FormatBool` |
| `Unquote` | `Quote` |

:::caution[Pitfalls]
- Always check the error from `Parse*`/`Atoi`; ignoring it turns bad input into 0.
- `ParseInt` with `bitSize 32` still returns `int64`; convert after checking.
:::
