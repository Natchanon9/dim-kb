---
title: "Goroutine Lifecycle"
description: "How goroutines start, block and end, and why main must wait"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A goroutine runs until its function returns. It can't be killed from outside; it has to be told to stop and then return by itself.

## Waiting for completion

```go title="lifecycle.go"
done := make(chan struct{})
go func() {
	defer close(done) // signal: finished
	fmt.Println("working")
}()
<-done // without this, main may return first
fmt.Println("finished")
```

## States

| State | Meaning |
| --- | --- |
| Runnable | Waiting for a CPU |
| Running | Executing on a thread |
| Waiting | Blocked on a channel, lock, I/O, sleep or syscall |
| Dead | Function returned |

When `main` returns, the program exits and every other goroutine stops mid-work, with no deferred calls run.

:::caution[Pitfalls]
- Plan the exit of every goroutine before you start it: who tells it to stop (context, closed channel), and who waits for it.
:::
