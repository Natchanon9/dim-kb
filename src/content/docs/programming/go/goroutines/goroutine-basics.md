---
title: "Goroutine Basics"
description: "Start concurrent work with the go keyword"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A goroutine is a function running concurrently with the rest of the program. Put `go` in front of a call to start one; it returns immediately.

## Example

```go title="basics.go"
var wg sync.WaitGroup
results := make([]int, 3)
for i := range 3 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		results[i] = i * i // each goroutine writes its own index
	}()
}
wg.Wait()
fmt.Println(results) // [0 1 4]
```

## Key facts

- The `go` statement doesn't return a value. Send results over a channel or write to memory you own.
- You must **wait** for goroutines (WaitGroup, channel, errgroup); nothing waits for them automatically.
- Goroutines are cheap: they start with a small stack (2 KB today) that grows as needed.

:::caution[Pitfalls]
- Two goroutines writing the same variable without a lock is a data race. Run tests with `-race`.
- `go f(x)` evaluates `x` now, in the calling goroutine.
:::
