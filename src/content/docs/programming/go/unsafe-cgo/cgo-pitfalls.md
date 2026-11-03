---
title: "cgo Pitfalls"
description: "Call overhead, builds, deployment and safety"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

"cgo is not Go." It works, but it gives up several things that make Go pleasant.

## Call overhead

```go title="bench_test.go"
func BenchmarkGoAdd(b *testing.B)  { /* plain Go add */ }
func BenchmarkCgoAdd(b *testing.B) { /* Add via C */ }
// GoAdd   ~0.2 ns/op
// CgoAdd  ~18 ns/op  (each call crosses the Go/C boundary)
```

Batch work so you cross the boundary rarely.

## Costs

| Area | Problem |
| --- | --- |
| Builds | Need a C toolchain; slower compiles |
| Cross-compiling | `GOOS=linux go build` no longer just works |
| Deployment | Binary may depend on shared libs (no `scratch` image) |
| Safety | C bugs crash the whole process; no race detector inside C |
| Tooling | Profilers and debuggers see less |

## Alternatives

Pure-Go libraries (e.g. `modernc.org/sqlite` instead of `mattn/go-sqlite3`), a subprocess, or a network service around the C code.

:::caution[Pitfalls]
- Blocking C calls tie up an OS thread each; thousands of concurrent ones create thousands of threads.
:::
