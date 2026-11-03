---
title: "Profiling with pprof"
description: "Find where CPU time and memory go"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

pprof samples a running program and shows which functions use CPU, allocate memory, or block. Profile before optimising.

## In a server

```go title="main.go"
import _ "net/http/pprof" // registers /debug/pprof/ on http.DefaultServeMux

go http.ListenAndServe("localhost:6060", nil) // keep it off the public port
```

```sh
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30  # CPU
go tool pprof http://localhost:6060/debug/pprof/heap                          # memory
```

## In code or tests

```go title="cpu.go"
f, _ := os.Create("cpu.out")
pprof.StartCPUProfile(f)
busyWork()
pprof.StopCPUProfile()
```

```sh
go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof -http=:8081 cpu.out
```

| Profile | Shows |
| --- | --- |
| profile (CPU) | Where time is spent |
| heap / allocs | Live memory / allocation sites |
| goroutine | All goroutine stacks |
| block / mutex | Waiting on channels / locks |

:::caution[Pitfalls]
- Never expose `/debug/pprof` publicly; it leaks internals and can be used for DoS.
:::
