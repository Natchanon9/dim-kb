---
title: "Fan-Out"
description: "Several workers read from one channel to share the work"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

Fan-out starts several goroutines that read from the same channel. Each value goes to exactly one worker, so the work spreads across them.

## Example

```go title="fanout.go"
jobs := emit(1, 2, 3, 4)
results := make(chan int)
var wg sync.WaitGroup
for range 3 { // 3 workers
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := range jobs {
			results <- j * 10
		}
	}()
}
go func() { wg.Wait(); close(results) }()
// collect results → [10 20 30 40] after sorting
```

Fan-out + fan-in together form the classic parallel processing shape; a fixed number of workers is a worker pool.

:::caution[Pitfalls]
- Results arrive in any order; carry an index or ID if you need to match them back.
:::
