---
title: "Pointer to Pointer"
description: "**T lets a function replace the caller's pointer"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`**T` is a pointer to a pointer. You need it only when a function must change which address the caller's pointer holds.

## Example

```go title="pp.go"
func setNil(pp **int) { *pp = nil }

x := 1
p := &x
pp := &p
**pp = 2
fmt.Println(x, *pp == p) // 2 true
setNil(&p)
fmt.Println(p == nil) // true
```

In everyday Go this is rare: returning the new pointer (`p = rebuild(p)`) is clearer.

:::caution[Pitfalls]
- If you reach for `**T`, check whether returning a value would be simpler.
:::
