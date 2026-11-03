---
title: "Value vs Pointer Receivers"
description: "When a method gets a copy and when it can modify the value"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A value receiver gets a copy. A pointer receiver gets the address, so it can modify the original and avoids copying.

## Example

```go title="counter.go"
type Counter struct{ n int }

func (c Counter) Value() int  { return c.n }
func (c *Counter) Inc()       { c.n++ } // changes the caller's Counter
func (c Counter) IncOnCopy()  { c.n++ } // changes a copy: no effect

var c Counter
c.Inc() // Go takes &c for you
c.Inc()
c.IncOnCopy()
fmt.Println(c.Value()) // 2
```

## Which to choose

| Use a pointer receiver when | Use a value receiver when |
| --- | --- |
| The method modifies the receiver | The type is small and immutable (`time.Time`, `Point`) |
| The struct is large | The type is a map, func or chan (already reference-like) |
| The struct holds a `sync.Mutex` | |

Be consistent: if any method needs a pointer receiver, give them all pointer receivers.

:::caution[Pitfalls]
- Copying a struct that contains a `sync.Mutex` copies the lock; `go vet` reports it.
:::
