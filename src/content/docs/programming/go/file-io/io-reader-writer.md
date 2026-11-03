---
title: "io.Reader & io.Writer"
description: "The two interfaces every Go I/O API is built on"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`io.Reader` (`Read(p []byte) (n int, err error)`) and `io.Writer` (`Write(p []byte) (n int, err error)`) connect files, network connections, HTTP bodies, compressors and hashes. Write code against them and it works with all of those.

## Your own Writer, composed readers

```go title="io.go"
type upperWriter struct{ w io.Writer }

func (u upperWriter) Write(p []byte) (int, error) {
	return u.w.Write(bytes.ToUpper(p))
}

var sb strings.Builder
io.Copy(upperWriter{&sb}, strings.NewReader("go io")) // GO IO

r := io.LimitReader(strings.NewReader("abcdef"), 3)
b, _ := io.ReadAll(r) // abc
```

## Useful helpers

| Helper | Does |
| --- | --- |
| `io.Copy(dst, src)` | Stream everything from src to dst |
| `io.ReadAll(r)` | Read to EOF into memory |
| `io.MultiWriter(a, b)` | Write to several writers |
| `io.TeeReader(r, w)` | Copy what is read into w |
| `io.LimitReader(r, n)` | Stop after n bytes |
| `io.Discard` | A writer that drops everything |

:::caution[Pitfalls]
- `Read` may return fewer bytes than asked, and data together with `io.EOF`. Use `io.ReadFull` when you need exactly n bytes.
:::
