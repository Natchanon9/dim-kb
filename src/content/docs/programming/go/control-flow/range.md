---
title: "Range"
description: "Iterate slices, strings, maps, channels and integers with range"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`for ... range` walks over a collection and gives you an index (or key) and a value on each iteration.

## What range returns

| Over | First value | Second value |
| --- | --- | --- |
| Slice / array | index | element copy |
| String | byte offset | rune (decoded UTF-8) |
| Map | key | value (random order) |
| Channel | element | (none) |
| Integer `n` (1.22+) | 0 … n-1 | (none) |

```go title="range.go"
for i, v := range []string{"a", "b"} {
	fmt.Println(i, v) // 0 a, 1 b
}
for i, r := range "gö" {
	fmt.Println(i, string(r)) // 0 g, 1 ö
}
```

Drop what you don't need: `for _, v := range s`, `for i := range s`, `for range ticker.C`.

:::caution[Pitfalls]
- `v` is a copy. `v.Count++` on a slice of structs changes nothing; use `s[i].Count++`.
- Map iteration order is random by design; sort the keys if order matters.
:::
