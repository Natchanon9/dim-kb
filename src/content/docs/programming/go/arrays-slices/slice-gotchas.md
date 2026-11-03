---
title: "Slice Gotchas"
description: "Shared backing arrays, append surprises and copies in range"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Most slice bugs come from two slices sharing one backing array without you noticing.

## Two appends, one slot

```go title="gotcha.go"
a := make([]int, 3, 4)
b := append(a, 1) // fits in cap: writes a's backing array
c := append(a, 2) // same slot again
fmt.Println(b[3], c[3]) // 2 2  (b changed too!)

safe := append(a[:3:3], 1) // full slice expression caps it, forcing a copy
fmt.Println(safe[3], b[3]) // 1 2
```

`s[low:high:max]` sets capacity to `max-low`, so the next `append` must allocate.

## Checklist

- **Function returns a sub-slice** of its input: callers can corrupt each other's data. Return `slices.Clone`.
- **`range` gives copies:** modify `s[i]`, not the loop value.
- **Tiny slice of a huge array** keeps the huge array in memory.
- **Appending to a slice you received** may or may not change the caller's data, depending on capacity.

:::caution[Pitfalls]
- These bugs appear only sometimes (they depend on capacity), which makes them hard to reproduce. Copy when ownership is unclear.
:::
