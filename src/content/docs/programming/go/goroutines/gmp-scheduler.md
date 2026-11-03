---
title: "GMP Scheduler"
description: "How the runtime maps goroutines onto OS threads"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

The Go runtime schedules goroutines itself, multiplexing many goroutines onto a few OS threads. The model has three parts.

| Letter | Is | Count |
| --- | --- | --- |
| **G** | Goroutine | Thousands to millions |
| **M** | Machine: an OS thread | Grows when threads block in syscalls |
| **P** | Processor: the right to run Go code, with a local run queue | `GOMAXPROCS` |

## How it runs

- An M must hold a P to run Gs. Each P has a queue of runnable Gs.
- An idle P **steals** work from other Ps' queues.
- When a G blocks on a syscall, its M is detached and the P moves to another M, so other Gs keep running.
- Network waits go to the netpoller instead of blocking a thread.
- Since Go 1.14, long-running Gs are **preempted** asynchronously, so one loop can't starve the others.

```go title="inspect.go"
fmt.Println("CPUs:", runtime.NumCPU())
fmt.Println("Ps (GOMAXPROCS):", runtime.GOMAXPROCS(0))
fmt.Println("goroutines:", runtime.NumGoroutine())
```

:::caution[Pitfalls]
- You rarely tune the scheduler. If performance looks off, profile first (`go tool trace`, pprof).
:::
