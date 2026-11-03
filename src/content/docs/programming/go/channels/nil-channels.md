---
title: "Nil Channels"
description: "Operations on a nil channel block forever, which can disable select cases"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Sending to or receiving from a `nil` channel blocks forever; closing it panics. In a `select`, that blocking is useful: setting a channel to nil disables its case.

## Merge two streams until both close

```go title="nilcase.go"
a, b := emit(1), emit(2)
sum := 0
for a != nil || b != nil {
	select {
	case v, ok := <-a:
		if !ok {
			a = nil // this case never fires again
			continue
		}
		sum += v
	case v, ok := <-b:
		if !ok {
			b = nil
			continue
		}
		sum += v
	}
}
fmt.Println(sum) // 3
```

Without the nil trick, a closed channel would be "ready" forever and the loop would spin.

:::caution[Pitfalls]
- A forgotten `make` leaves a channel nil; the goroutine using it hangs silently.
:::
