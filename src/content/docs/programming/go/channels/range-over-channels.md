---
title: "Range over Channels"
description: "Receive until the channel is closed"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`for v := range ch` receives until the channel is closed **and** drained. It is the normal way to consume a stream.

## Example

```go title="range.go"
ch := make(chan string, 3)
ch <- "a"
ch <- "b"
close(ch)
for s := range ch {
	fmt.Println(s) // a, b
}
v, ok := <-ch // closed and empty
fmt.Printf("%q %v\n", v, ok) // "" false
```

## Closing rules

- Close to say "no more values", not to free resources.
- The **sender** closes, and only once.
- With several senders, close after all of them finish (WaitGroup + one closer goroutine).

:::caution[Pitfalls]
- `range` over a channel nobody closes never ends: a goroutine leak.
:::
