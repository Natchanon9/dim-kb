---
title: "GOMAXPROCS"
description: "How many threads run Go code at once"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`GOMAXPROCS` is the number of Ps: how many goroutines can execute Go code **in parallel**. It defaults to the number of CPUs available.

## Reading and setting

```go title="procs.go"
prev := runtime.GOMAXPROCS(2) // set; returns the previous value
defer runtime.GOMAXPROCS(prev)
runtime.GOMAXPROCS(0) // 0 = query only → 2
```

Or set the `GOMAXPROCS=4` environment variable.

## Containers

Since Go 1.25 the default respects a container's CPU limit (cgroups) and updates if it changes. On older versions in Kubernetes, the default was the node's CPU count, which caused throttling; teams used `go.uber.org/automaxprocs`.

:::caution[Pitfalls]
- Raising it above the CPU count doesn't add parallelism; it adds contention.
- Goroutines blocked in I/O don't count against it, so I/O-heavy servers don't need a higher value.
:::
