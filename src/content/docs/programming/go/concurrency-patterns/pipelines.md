---
title: "Pipelines"
description: "Chain stages connected by channels"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

A pipeline is a series of stages. Each stage receives from an input channel, does one job, and sends to an output channel it owns and closes.

## Example

```go title="pipeline.go"
func source(nums ...int) <-chan int { /* send nums, then close */ }

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func sum(in <-chan int) int { /* add everything from in */ }

fmt.Println(sum(square(source(1, 2, 3)))) // 14
```

Stages run concurrently: while `sum` adds one value, `square` works on the next. Slow stages can be fanned out to several goroutines.

:::caution[Pitfalls]
- If the consumer stops early, upstream stages block forever. Pass a `ctx` and `select` on `ctx.Done()` in every send.
:::
