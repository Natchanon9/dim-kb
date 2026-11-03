---
title: "Interface Composition"
description: "Build bigger interfaces by embedding small ones"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

An interface can embed other interfaces. The result requires all their methods. This is how `io.ReadWriter` is built from `io.Reader` and `io.Writer`.

## Example

```go title="compose.go"
type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface {
	Reader
	Writer
}

var rw ReadWriter = &memFile{} // memFile has Read and Write
rw.Write("hi")
fmt.Println(rw.Read()) // hi
```

## Design rule

Start with one- or two-method interfaces and compose them where a function needs more. Functions should ask for the smallest interface they use.

:::caution[Pitfalls]
- Big interfaces (10+ methods) are hard to implement and fake in tests; split them.
:::
