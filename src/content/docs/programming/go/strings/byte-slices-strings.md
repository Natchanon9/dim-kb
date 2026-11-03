---
title: "Byte Slices & Strings"
description: "Converting between []byte and string, and the bytes package"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`[]byte` is the mutable counterpart of `string`. I/O, hashing and JSON work in `[]byte`; most text APIs work in `string`.

## Converting

```go title="bytes.go"
b := []byte("hello") // copies the string's bytes
b[0] = 'H'           // []byte is mutable
fmt.Println(string(b), b, string(bytes.ToUpper(b)))
// Hello [72 101 108 108 111] HELLO
```

Each conversion copies the data, because strings must stay immutable.

## bytes mirrors strings

`bytes` has the same helpers as `strings` (`Contains`, `Split`, `TrimSpace`, `Equal` …) plus `bytes.Buffer`, a growable buffer that is both an `io.Reader` and `io.Writer`.

:::caution[Pitfalls]
- Converting back and forth in a loop allocates every time; pick one type and stay in it.
- Compare byte slices with `bytes.Equal`, not `==` (which doesn't compile for slices).
:::
