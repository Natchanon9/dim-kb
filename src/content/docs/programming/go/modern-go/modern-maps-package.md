---
title: "maps Package (Modern)"
description: "Keys, Values, All and Collect (Go 1.23)"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Go 1.23 added iterator functions to `maps`. The most used one-liner: sorted keys.

## Example

```go title="maps.go"
m := map[string]int{"b": 2, "a": 1}
slices.Sorted(maps.Keys(m))        // [a b]
copyOfM := maps.Collect(maps.All(m)) // iterator → new map
```

| Function | Does |
| --- | --- |
| `Keys(m)` | Iterator of keys |
| `Values(m)` | Iterator of values |
| `All(m)` | Iterator of key, value |
| `Collect(seq2)` | Key-value iterator → new map |
| `Insert(m, seq2)` | Add pairs into an existing map |

The 1.21 functions (`Clone`, `Equal`, `DeleteFunc`, `Copy`) are on [maps Package](/dim-kb/programming/go/maps/maps-package/).

:::caution[Pitfalls]
- `maps.Keys` returns an iterator, not a slice, in the standard library; old code using `golang.org/x/exp/maps.Keys` got a slice.
:::
