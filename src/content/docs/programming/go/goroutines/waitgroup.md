---
title: "WaitGroup"
description: "Wait for a group of goroutines to finish"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`sync.WaitGroup` counts running goroutines. `Add` before starting, `Done` when each finishes, `Wait` blocks until the count is zero.

## Example

```go title="wg.go"
var (
	wg    sync.WaitGroup
	mu    sync.Mutex
	total int
)
for _, n := range []int{1, 2, 3} {
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		total += n
		mu.Unlock()
	}()
}
wg.Wait()
fmt.Println(total) // 6
```

Go 1.25 adds `wg.Go(func() { ... })`, which does `Add(1)`, `go` and `Done` for you.

:::caution[Pitfalls]
- Call `Add` **before** `go`, not inside the goroutine, or `Wait` may return too early.
- Pass a WaitGroup by pointer; copying it breaks the counter (`go vet` warns).
- A WaitGroup doesn't collect errors; use `errgroup` for that.
:::
