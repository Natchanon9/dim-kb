---
title: "Sizeof, Alignof, Offsetof"
description: "Struct layout, padding and field order"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Fields are aligned to their size, so the compiler inserts padding. Field order changes a struct's size.

## Example

```go title="layout.go"
type padded struct {
	A bool  // 1 + 7 padding
	B int64 // 8
	C bool  // 1 + 7 padding
}

type packed struct {
	B int64 // 8
	A bool  // 1
	C bool  // 1 + 6 padding
}

unsafe.Sizeof(padded{}), unsafe.Sizeof(packed{}) // 24 16
unsafe.Alignof(int64(0))                         // 8
unsafe.Offsetof(padded{}.B)                      // 8
```

Ordering fields from largest to smallest usually minimises padding. It matters for structs stored by the million (cache entries, slices of records).

The `fieldalignment` analyzer (golang.org/x/tools) reports structs that could be smaller.

:::caution[Pitfalls]
- Don't reorder for size at the cost of readability unless profiling shows memory pressure.
:::
