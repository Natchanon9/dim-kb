---
title: "Reflection Performance"
description: "How much reflection costs and how to reduce it"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Reflection is much slower than direct code: it allocates, does lookups by name, and defeats compiler optimisations.

## Benchmark

```go title="bench_test.go"
func BenchmarkDirectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = p.X
	}
}

func BenchmarkReflectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = int(reflect.ValueOf(p).FieldByName("X").Int())
	}
}
// DirectField   ~0.2 ns/op
// ReflectField  ~35 ns/op   (about 100x slower on this machine)
```

## Making it cheaper

- **Cache** per-type work (field indexes, parsed tags) in a `sync.Map` keyed by `reflect.Type`; that's what `encoding/json` does.
- Use `Field(i)` with a cached index instead of `FieldByName`.
- Generate code instead (`go generate`, `sqlc`, `easyjson`) for hot paths.

:::caution[Pitfalls]
- Measure before optimising; for a request that does I/O, 35 ns rarely matters.
:::
