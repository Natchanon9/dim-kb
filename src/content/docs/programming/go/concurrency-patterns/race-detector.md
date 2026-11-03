---
title: "Race Detector"
description: "Find data races with -race"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

A data race is two goroutines touching the same memory at the same time, with at least one write and no synchronisation. Results become unpredictable. The race detector finds them at runtime.

## A race and its fix

```go title="race.go"
count := 0
for range 2 {
	go func() { count++ }() // RACE: unsynchronised read-modify-write
}

var safe atomic.Int64
for range 2 {
	go func() { safe.Add(1) }() // fixed
}
```

```sh
go test -race ./...
go run -race .
```

The report shows both stack traces: where the write happened and where the conflicting access happened.

## Use it

- Always in CI for tests.
- In staging builds if you can afford ~5–10x CPU and memory overhead.

:::caution[Pitfalls]
- It only finds races on paths that run. No report doesn't prove there are none.
:::
