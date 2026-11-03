---
title: "unsafe.Pointer"
description: "Convert pointers between types and zero-copy conversions"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`unsafe.Pointer` can be converted to and from any pointer type, bypassing the type system. The spec lists the few patterns that are valid.

## Example

```go title="pointer.go"
b := []byte("hello")
s := unsafe.String(unsafe.SliceData(b), len(b)) // no copy; b must never change again
fmt.Println(s) // hello

x := int32(7)
p := unsafe.Pointer(&x)
fmt.Println(*(*int32)(p)) // 7
```

## Rules that keep it valid

- Convert `*T1 → unsafe.Pointer → *T2` only when the memory layouts are compatible.
- Don't store a pointer as `uintptr` across statements; the GC doesn't see it and the object may move or be freed.
- The zero-copy string trick is only safe if the bytes are never modified afterwards.

`go vet` checks some misuses ("possible misuse of unsafe.Pointer").

:::caution[Pitfalls]
- Bugs here show up as random crashes or corrupted data far from the cause.
:::
