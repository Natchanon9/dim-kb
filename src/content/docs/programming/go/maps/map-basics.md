---
title: "Map Basics"
description: "Key-value lookup with map[K]V"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A map stores key-value pairs with fast lookup by key. Keys must be comparable with `==` (strings, numbers, structs of those, arrays); slices, maps and funcs can't be keys.

## Create, write, read

```go title="map.go"
ages := map[string]int{"ann": 30}
ages["bob"] = 25
fmt.Println(ages["ann"], ages["zed"], len(ages)) // 30 0 2
```

Reading a missing key returns the value type's zero value, not an error.

## Nil map

```go title="nil.go"
var m map[string]int // nil map: reads work, writes panic
fmt.Println(m["x"], m == nil) // 0 true
```

Create maps with a literal or `make(map[K]V)` before writing.

:::caution[Pitfalls]
- `m["k"] = v` on a nil map panics: "assignment to entry in nil map". Common with map fields in structs.
- Maps aren't safe for concurrent writes; the runtime may crash with "concurrent map writes".
:::
