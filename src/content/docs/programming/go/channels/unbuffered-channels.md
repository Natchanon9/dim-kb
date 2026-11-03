---
title: "Unbuffered Channels"
description: "A send and a receive meet: synchronous handoff"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

An unbuffered channel (`make(chan T)`) has no queue. A send blocks until a receiver takes the value, so both goroutines meet at that moment.

## Example

```go title="handoff.go"
ch := make(chan int)
done := make(chan struct{})
go func() {
	v := <-ch
	fmt.Println("got", v) // got 42
	close(done)
}()
ch <- 42 // blocks until the goroutine receives
<-done
```

## Why it's useful

The handoff is a guarantee: when the send returns, the receiver has the value. That makes unbuffered channels good for signalling and for passing ownership of data.

| | Unbuffered | Buffered |
| --- | --- | --- |
| Send blocks until | A receiver takes it | There is space |
| Guarantees delivery at send time | Yes | No |

:::caution[Pitfalls]
- Sending and receiving on the same unbuffered channel in one goroutine deadlocks.
:::
