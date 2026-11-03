---
title: "String Formatting"
description: "Sprintf verbs for numbers, width, padding and alignment"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`fmt.Sprintf` builds a formatted string with the same verbs as `Printf`. Width and flags go between `%` and the verb.

## Verbs and flags

```go title="format.go"
fmt.Sprintf("%05d|%x|%X|%o|%b", 42, 255, 255, 8, 5)
// 00042|ff|FF|10|101
fmt.Sprintf("|%8.3f|%-8s|%8s|", 3.14159, "left", "right")
// |   3.142|left    |   right|
fmt.Printf("%v %d%%\n", []int{1, 2}, 50)
// [1 2] 50%
```

| Flag | Meaning |
| --- | --- |
| `%8s` | Width 8, right-aligned |
| `%-8s` | Left-aligned |
| `%05d` | Pad with zeros |
| `%.3f` | 3 decimal places |
| `%x` / `%X` | Hex (also works on strings and `[]byte`) |
| `%%` | A literal `%` |

Types with a `String() string` method print through it with `%v` and `%s`.

:::caution[Pitfalls]
- Wrong verb/argument pairs print `%!d(string=hi)` at runtime; `go vet` catches most of them.
- Calling `fmt.Sprintf("%v", x)` inside `x.String()` recurses forever.
:::
