---
title: "Done Channel"
description: "Close a channel to tell goroutines to stop"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Closing a channel wakes **every** receiver at once. A `done` channel uses that as a broadcast "stop" signal.

## Example

```go title="done.go"
func numbers(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

done := make(chan struct{})
nums := numbers(done)
fmt.Println(<-nums, <-nums) // 0 1
close(done)                 // the goroutine returns
```

`chan struct{}` carries no data, only the signal.

## Today: use context

`context.Context` is the standard version of this pattern (`ctx.Done()` is a done channel), with deadlines and propagation across APIs. Use it in new code.

:::caution[Pitfalls]
- Send a value on `done` and only one goroutine sees it; close it to reach all of them.
:::
