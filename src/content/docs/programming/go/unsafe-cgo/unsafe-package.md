---
title: "unsafe Package"
description: "Step outside Go's type safety, and why you rarely should"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`unsafe` gives low-level access to memory layout and pointer conversion. Code that imports it loses Go's memory-safety and compatibility guarantees.

## What's in it

```go title="sizes.go"
fmt.Println(unsafe.Sizeof(int64(0)), unsafe.Sizeof(""), unsafe.Sizeof([]int{}))
// 8 16 24  (64-bit: string = ptr+len, slice = ptr+len+cap)
```

| Item | Does |
| --- | --- |
| `unsafe.Pointer` | Convert between any pointer types |
| `Sizeof`, `Alignof`, `Offsetof` | Memory layout, at compile time |
| `String`, `StringData`, `Slice`, `SliceData` (1.20+) | Zero-copy string/slice views |
| `Add` | Pointer arithmetic |

## When it's justified

Performance-critical libraries (serialization, hashing), syscalls, and interop with C. Application code almost never needs it.

:::caution[Pitfalls]
- Code using `unsafe` may break on a new Go version or another architecture; the Go 1 compatibility promise doesn't cover it.
:::
