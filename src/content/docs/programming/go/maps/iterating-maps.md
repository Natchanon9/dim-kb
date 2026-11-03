---
title: "Iterating Maps"
description: "range over a map and get a stable order"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`for k, v := range m` visits every entry, in a **random** order that changes between runs on purpose.

## Sorted iteration

```go title="sorted.go"
m := map[string]int{"b": 2, "a": 1, "c": 3}
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
slices.Sort(keys)
for _, k := range keys {
	fmt.Println(k, m[k]) // a 1, b 2, c 3
}
```

From Go 1.23 this is one line: `slices.Sorted(maps.Keys(m))`.

## Changing a map while ranging

Deleting entries during `range` is safe. Entries added during `range` may or may not be visited.

:::caution[Pitfalls]
- Never rely on map order in output, tests or JSON arrays you build from a map.
:::
