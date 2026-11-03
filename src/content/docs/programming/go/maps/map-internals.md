---
title: "Map Internals"
description: "What a map value is, growth, and why order is random"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

A map value is a pointer to a runtime hash table. Copying or passing a map shares the same table.

## Passing maps

```go title="share.go"
func addItem(m map[string]int) { m["x"] = 1 } // no pointer needed

m := make(map[string]int, 100) // size hint avoids regrowth
addItem(m)
fmt.Println(m) // map[x:1]
```

## What happens inside

- Keys are hashed into buckets; lookups are O(1) on average.
- When the table gets too full, it grows and entries are moved. That is why `&m[k]` is forbidden: elements can move.
- Iteration starts at a random position so code can't depend on order.
- Since Go 1.24 the runtime uses a Swiss-table design, which is faster but behaves the same from your code.

:::caution[Pitfalls]
- Maps never shrink after `delete`. A map that once held millions of keys keeps that memory; rebuild it if needed.
:::
