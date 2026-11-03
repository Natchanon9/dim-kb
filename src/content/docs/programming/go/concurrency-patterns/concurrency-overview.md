---
title: "Concurrency Overview"
description: "Concurrency vs parallelism and which tool to reach for"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Concurrency is structuring a program as independent tasks; parallelism is running tasks at the same instant on several CPUs. Go gives you concurrency, and the runtime adds parallelism when CPUs are available.

## Choosing a tool

| Problem | Tool |
| --- | --- |
| Run N tasks and wait | `sync.WaitGroup` |
| Run N tasks, stop on first error | `errgroup` |
| Pass data or ownership between goroutines | Channels |
| Protect shared state | `sync.Mutex` / `RWMutex` |
| A single counter or flag | `sync/atomic` |
| Initialise once | `sync.Once`, `sync.OnceValue` |
| Cancel and set deadlines | `context` |
| Limit concurrency | Worker pool or semaphore channel |
| Reuse temporary objects | `sync.Pool` |

## Channels or mutex?

Use channels when goroutines hand work to each other (pipelines, pools, events). Use a mutex when several goroutines touch one piece of state (a cache, a counter map). Both are idiomatic; pick the one that makes the code simpler.

:::caution[Pitfalls]
- Concurrency adds bugs. Make a sequential version work first, then add goroutines where profiling shows waiting.
:::
