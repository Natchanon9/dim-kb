---
title: "strings Package"
description: "Search, split, join, trim and replace text"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`strings` holds the everyday text helpers. All functions return new strings.

## Common functions

```go title="strings.go"
t := strings.TrimSpace("  Hello, Go  ")
strings.ToUpper(t)                      // HELLO, GO
strings.Contains(t, "Go")               // true
strings.HasPrefix(t, "He")              // true
strings.Index(t, "Go")                  // 7
strings.Split("a,b,c", ",")             // [a b c]
strings.Join([]string{"x", "y"}, "-")   // x-y
strings.ReplaceAll("aaa", "a", "b")     // bbb
strings.Fields(" a  b ")                // [a b]
strings.EqualFold("Go", "GO")           // true (case-insensitive)
key, value, found := strings.Cut("key=value", "=") // key value true
```

`Cut` (Go 1.18+) replaces most `Index` + slicing code for "split once".

:::caution[Pitfalls]
- `Split("", ",")` returns `[""]` (one empty element), not an empty slice.
- Comparing with `ToLower(a) == ToLower(b)` allocates twice; use `EqualFold`.
:::
