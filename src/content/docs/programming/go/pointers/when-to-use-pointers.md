---
title: "When to Use Pointers"
description: "Mutation, large structs, shared state and optional values"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Use pointers for a reason, not by habit. Values are simpler and often faster.

## Use a pointer when

- **The callee must modify** the value.
- **The struct is large** (many fields, big arrays) and copying it is costly.
- **The value must be shared** or holds a lock (`sync.Mutex`).
- **You need "absent"** distinct from the zero value:

```go title="patch.go"
// "not sent" (nil) differs from "set to zero"
type Patch struct {
	Name *string
	Age  *int
}

age := 0
patch := Patch{Age: &age}
fmt.Println(patch.Name == nil, *patch.Age) // true 0
```

## Prefer a value when

- Small, immutable data (`time.Time`, points, money amounts).
- Maps, slices, channels and funcs: they already refer to shared data.

:::caution[Pitfalls]
- `[]*T` vs `[]T`: pointers in a slice add allocations and GC work; use `[]T` unless elements must be shared or are huge.
:::
