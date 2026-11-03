---
title: "Goroutines vs Threads"
description: "Why goroutines are cheaper than OS threads"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Goroutines are scheduled by the Go runtime, not the OS. That makes them far cheaper, so one goroutine per request or per task is normal.

## Comparison

| | Goroutine | OS thread |
| --- | --- | --- |
| Initial stack | ~2 KB, grows and shrinks | ~1–8 MB, fixed |
| Creation | Microseconds, no syscall | Syscall, much slower |
| Switch | In user space, cheap | Kernel context switch |
| Identity | No ID exposed | Thread ID |
| Typical count | 100k+ | Thousands |

```go title="many.go"
var (
	wg sync.WaitGroup
	n  atomic.Int64
)
for range 10_000 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.Add(1)
	}()
}
wg.Wait()
fmt.Println(n.Load()) // 10000
```

:::caution[Pitfalls]
- Cheap isn't free: unbounded goroutines hitting a database or API will overwhelm it. Bound concurrency with a worker pool or semaphore.
:::
