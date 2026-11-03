---
title: "Channel Directions"
description: "Send-only and receive-only channel types"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

A function parameter can restrict a channel to one direction. The compiler then rejects the wrong operation.

## Example

```go title="directions.go"
func produce(out chan<- int) { // send-only
	for i := range 3 {
		out <- i
	}
	close(out)
}

func consume(in <-chan int) int { // receive-only
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}

ch := make(chan int)
go produce(ch)
fmt.Println(consume(ch)) // 3
```

| Type | Can |
| --- | --- |
| `chan T` | Send and receive |
| `chan<- T` | Send (and close) only |
| `<-chan T` | Receive only |

A bidirectional channel converts implicitly to either direction.

:::caution[Pitfalls]
- Closing a receive-only channel doesn't compile, which is the point: only senders close.
:::
