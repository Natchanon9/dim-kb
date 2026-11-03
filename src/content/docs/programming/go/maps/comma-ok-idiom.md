---
title: "Comma-ok Idiom"
description: "Tell a missing key from a zero value"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`v, ok := m[k]` returns a second boolean that says whether the key exists. Use it whenever the zero value is a valid value.

## Example

```go title="ok.go"
stock := map[string]int{"apple": 0}
v, ok := stock["apple"] // 0 true  → present, value is 0
v, ok = stock["pear"]   // 0 false → absent
```

Usually combined with if-init:

```go
if price, ok := prices[item]; ok {
	total += price
}
```

The same `v, ok` shape appears in type assertions (`x.(T)`) and channel receives (`v, ok := <-ch`).

:::caution[Pitfalls]
- Without `ok`, "stock is 0" and "no such product" look identical.
:::
