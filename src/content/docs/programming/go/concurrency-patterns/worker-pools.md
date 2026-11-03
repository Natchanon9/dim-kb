---
title: "Worker Pools"
description: "A fixed number of goroutines processing a queue of jobs"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A worker pool starts N goroutines that pull jobs from one channel. It caps concurrency, which protects databases and APIs from overload.

## Example

```go title="pool.go"
func squareAll(jobs []int, workers int) []int {
	in := make(chan int)
	out := make(chan int, len(jobs))
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range in {
				out <- j * j
			}
		}()
	}
	for _, j := range jobs {
		in <- j
	}
	close(in) // workers leave their range loop
	wg.Wait()
	close(out)
	// collect out → [1 4 9 16]
}
```

## Sizing

- CPU-bound work: about `runtime.NumCPU()` workers.
- I/O-bound work: more, limited by what the downstream service can take.

:::caution[Pitfalls]
- Close the jobs channel exactly once, after the last send; close results only after all workers finish.
:::
