---
title: "cgo Basics"
description: "Call C code from Go with import \"C\""
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

cgo lets Go call C. C code goes in a comment right above `import "C"`, and its functions appear under the `C.` prefix.

## Example

```go title="cgo.go"
//go:build cgo

package mathc

/*
static int add(int a, int b) { return a + b; }
*/
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
// Add(2, 3) → 5
```

## Build requirements

- A C compiler (gcc/clang) on the build machine.
- `CGO_ENABLED=1` (default when a compiler is found; disabled for cross-compiling).
- The `//go:build cgo` tag lets the package build without cgo when you provide a fallback.

:::caution[Pitfalls]
- `import "C"` must directly follow the C comment, with no blank line.
- cgo isn't allowed in `_test.go` files; put it in a normal file and test that.
:::
