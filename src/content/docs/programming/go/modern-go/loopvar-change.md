---
title: "Loop Variable Change"
description: "Each loop iteration gets its own variable since Go 1.22"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Before Go 1.22, a `for` loop had one variable shared by all iterations, so closures and goroutines captured the last value. Since 1.22, each iteration has its own copy.

## Example

```go title="loopvar.go"
var fns []func() int
for i := range 3 {
	fns = append(fns, func() int { return i })
}
for _, f := range fns {
	fmt.Print(f()) // 012 on Go 1.22+ (was 333 with a 3-part loop before)
}
```

## What it fixes

- `go func() { use(v) }()` inside `range` used to see only the last element.
- Taking `&v` in a loop gave the same pointer every iteration.

The old workaround `v := v` is now unnecessary; tools like `copyloopvar` flag it.

:::caution[Pitfalls]
- The new behaviour depends on the `go` line in `go.mod`, per module. Old modules keep old semantics.
:::
