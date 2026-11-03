---
title: "gob Encoding"
description: "Go's binary format for Go-to-Go data"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`encoding/gob` is a compact, self-describing binary format for sending Go values between Go programs.

## Example

```go title="gob.go"
type Point struct{ X, Y int }

var buf bytes.Buffer
gob.NewEncoder(&buf).Encode(Point{1, 2})
var p Point
gob.NewDecoder(&buf).Decode(&p)
fmt.Println(p) // {1 2}
```

## When to use

Caches or files read only by your own Go code, and `net/rpc`. For anything another language reads, use JSON or Protocol Buffers.

:::caution[Pitfalls]
- Interface values need `gob.Register(ConcreteType{})` before encoding.
- Unexported fields are skipped.
:::
