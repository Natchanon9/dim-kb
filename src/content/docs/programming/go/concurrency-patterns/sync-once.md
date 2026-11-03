---
title: "sync.Once"
description: "Run initialisation exactly once, safely"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`sync.Once` runs a function exactly once, even if many goroutines call it at the same time. Others wait until the first call finishes.

## Example

```go title="once.go"
var once sync.Once
calls := 0
for range 3 {
	once.Do(func() { calls++ }) // runs only the first time
}

config := sync.OnceValue(func() string { // Go 1.21+
	return "loaded"
})
fmt.Println(calls, config(), config()) // 1 loaded loaded
```

`OnceValue` and `OnceValues` (value + error) cache the result, which is the usual lazy-singleton need.

:::caution[Pitfalls]
- If the function panics, `Once` still counts it as done; later calls do nothing.
- Calling `once.Do` inside the function it runs deadlocks.
:::
