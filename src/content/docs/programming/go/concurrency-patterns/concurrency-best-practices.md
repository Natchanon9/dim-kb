---
title: "Concurrency Best Practices"
description: "Rules that keep concurrent Go code correct"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Most concurrency bugs come from unclear ownership and goroutines without an exit.

## Checklist

- **Know how every goroutine ends** before starting it.
- **Pass `context.Context`** as the first parameter of anything that blocks, and respect `ctx.Done()`.
- **The sender closes** a channel; receivers never close.
- **Prefer synchronous APIs.** Let the caller add concurrency; don't start goroutines inside a library function and return.
- **Bound concurrency** with a pool, semaphore or `errgroup.SetLimit`.
- **Keep locks next to the data** and critical sections short.
- **Run `-race`** in CI.
- **Measure before optimising:** extra goroutines don't help CPU-bound code beyond the number of cores.

:::caution[Pitfalls]
- `time.Sleep` to "wait for a goroutine" is a bug; use a WaitGroup or channel.
:::
