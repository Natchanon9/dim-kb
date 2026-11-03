---
title: "Goroutine Debugging"
description: "Race detector, stack dumps and goroutine profiles"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

Concurrent bugs are hard to reproduce, so lean on tools that observe the running program.

## Tools

| Tool | Finds |
| --- | --- |
| `go test -race`, `go run -race` | Data races |
| `kill -QUIT <pid>` / Ctrl+\ | Dumps every goroutine's stack and exits |
| `runtime.Stack(buf, true)` | Stack dump from inside the program |
| pprof `/debug/pprof/goroutine?debug=2` | Live goroutine stacks via HTTP |
| `go tool trace` | Scheduling, blocking and GC timeline |
| Delve (`dlv`) | Step through goroutines in a debugger |

```go title="stack.go"
buf := make([]byte, 1<<16)
n := runtime.Stack(buf, true) // true: every goroutine
fmt.Println(string(buf[:n]))  // "goroutine 1 [running]: ..."
```

A deadlock where every goroutine is blocked ends with `fatal error: all goroutines are asleep - deadlock!` and a dump.

:::caution[Pitfalls]
- The race detector only finds races on code paths that actually run; test concurrent paths.
:::
