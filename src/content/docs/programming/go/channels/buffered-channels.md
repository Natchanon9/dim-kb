---
title: "Buffered Channels"
description: "Channels with capacity: sends don't wait until the buffer is full"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A buffered channel has a queue. Sends succeed immediately while there is room; receives succeed while it holds values.

## Example

```go title="buffered.go"
ch := make(chan int, 2)
ch <- 1 // no receiver needed while there is room
ch <- 2
fmt.Println(len(ch), cap(ch)) // 2 2
fmt.Println(<-ch, <-ch)       // 1 2 (FIFO)
```

## When to buffer

- **Size 1** so a goroutine can deliver a result and exit even if nobody reads it (avoids leaks on timeout).
- **Known batch size**, e.g. `make(chan Result, len(jobs))`.
- **Semaphore**: `make(chan struct{}, n)` limits concurrency to n.

Don't add a big buffer to "fix" a deadlock; it usually hides the bug until load rises.

:::caution[Pitfalls]
- `len(ch)` is a snapshot; it may change before you use it. Don't base logic on it.
:::
