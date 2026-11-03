---
title: "Channel Basics"
description: "Send and receive values between goroutines"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A channel is a typed pipe between goroutines. One side sends with `ch <- v`, the other receives with `<-ch`, and the operation synchronises them.

## Example

```go title="ping.go"
ch := make(chan string)
go func() { ch <- "ping" }() // send
msg := <-ch                  // receive: waits for the send
fmt.Println(msg)             // ping
```

## Operations

| Operation | Syntax | Notes |
| --- | --- | --- |
| Create | `make(chan T)` / `make(chan T, n)` | Zero value is `nil` |
| Send | `ch <- v` | Blocks until a receiver or buffer space |
| Receive | `v := <-ch` / `v, ok := <-ch` | `ok` is false when closed and empty |
| Close | `close(ch)` | Only the sender should close |

"Don't communicate by sharing memory; share memory by communicating." Pass ownership of data through channels instead of locking shared variables.

:::caution[Pitfalls]
- Sending on a closed channel panics; closing twice panics.
- Every goroutine blocked on a channel with no partner → `all goroutines are asleep - deadlock!`.
:::
