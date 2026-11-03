---
title: "Calling C from Go"
description: "Pass strings and memory across the boundary safely"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Go and C manage memory differently. Strings and buffers must be copied into C memory and freed by you.

## Passing a string

```go title="strlen.go"
/*
#include <stdlib.h>
#include <string.h>
static size_t c_strlen(const char *s) { return strlen(s); }
*/
import "C"

func CStrlen(s string) int {
	cs := C.CString(s) // malloc'd copy, not tracked by the Go GC
	defer C.free(unsafe.Pointer(cs))
	return int(C.c_strlen(cs))
}
// CStrlen("hello, C") → 8
```

| Conversion | Function |
| --- | --- |
| Go string → C | `C.CString` (free it!) |
| C string → Go | `C.GoString` |
| Go []byte → C | `C.CBytes` (free it!) |
| C buffer → Go | `C.GoBytes(ptr, n)` |
| Linking a library | `// #cgo LDFLAGS: -lsqlite3` |

:::caution[Pitfalls]
- C may not keep Go pointers after the call returns (cgo pointer rules); copy data instead.
- Every `C.CString`/`C.CBytes` without `C.free` is a memory leak.
:::
