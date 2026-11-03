---
title: "Fan-In"
description: "Merge several channels into one"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Fan-in combines values from several input channels into one output channel. One goroutine per input copies values; a final goroutine closes the output after all inputs are done.

## Example

```go title="merge.go"
func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}
	go func() { wg.Wait(); close(out) }()
	return out
}

for v := range merge(emit(1, 2), emit(3)) { ... } // 1, 2, 3 in some order
```

Use it to combine results from parallel workers or several event sources.

:::caution[Pitfalls]
- Output order is not preserved across inputs.
- Add a `ctx.Done()` case if consumers may stop early; otherwise the copy goroutines leak.
:::
