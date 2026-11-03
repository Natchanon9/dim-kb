---
title: "Benchmarks"
description: "Measure speed and allocations with BenchmarkXxx"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

A benchmark is `BenchmarkXxx(b *testing.B)`. The framework picks how many iterations to run to get a stable time per operation.

## Example

```go title="join_test.go"
func BenchmarkJoin(b *testing.B) {
	parts := []string{"a", "b", "c"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ { // Go 1.24+: for b.Loop() { ... }
		strings.Join(parts, ",")
	}
}
```

```sh
go test -bench=. -benchmem ./...
# BenchmarkJoin-16   12000000   95 ns/op   8 B/op   1 allocs/op
```

## Comparing changes

Run several times (`-count=10`) before and after, then compare with `benchstat old.txt new.txt`; single runs are noisy.

:::caution[Pitfalls]
- Setup inside the loop gets measured; do it before the loop or call `b.ResetTimer()`.
- The compiler may remove work whose result is unused; `b.Loop()` (1.24+) prevents that.
:::
