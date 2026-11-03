---
title: "sync.Pool"
description: "Reuse temporary objects to reduce allocations"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`sync.Pool` keeps spare objects for reuse, cutting allocations and GC work in hot paths like encoding buffers.

## Example

```go title="pool.go"
var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func render(name string) string {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset() // a pooled object may hold old data
	buf.WriteString("hello, ")
	buf.WriteString(name)
	return buf.String() // hello, go
}
```

## Rules

- The pool can drop objects at any GC; never use it as a cache.
- Always reset objects before use.
- Measure first: it only pays off for frequently allocated, short-lived objects.

:::caution[Pitfalls]
- Don't put very large buffers back; they pin memory. Check `buf.Cap()` and drop oversized ones.
:::
