---
title: "Blank Identifier"
description: "The _ placeholder for values, imports and compile-time checks"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

`_` is a write-only placeholder. Assigning to it discards the value, which keeps the compiler happy about unused results.

## Uses

```go title="blank.go"
n, _ := strconv.Atoi("42") // ignore the error on purpose
for _, v := range items {}  // ignore the index

import _ "image/png" // import only for its init side effect

var _ fmt.Stringer = celsius(0) // compile-time check: celsius implements Stringer
```

The last line fails to compile if `celsius` ever stops satisfying the interface.

:::caution[Pitfalls]
- Ignoring an error with `_` should be a deliberate, rare choice; reviewers will ask why.
- `_` can't be read: `x := _` doesn't compile.
:::
