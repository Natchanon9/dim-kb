---
title: "defer"
description: "Run cleanup when the function returns, in LIFO order"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`defer` schedules a call to run when the surrounding function returns, even if it panics. Use it to put cleanup right next to the setup.

## Order and argument evaluation

```go title="defer.go"
for i := range 3 {
	defer fmt.Print(" ", i) // last in, first out
}
fmt.Print("start") // prints: start 2 1 0

x := 1
defer fmt.Println("deferred:", x) // x evaluated now → 1
x = 2
```

## Typical cleanup

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close()
```

Same for `mu.Unlock()`, `resp.Body.Close()`, `rows.Close()`, `cancel()`.

:::caution[Pitfalls]
- `defer` in a loop runs only when the **function** returns. Thousands of deferred `Close` calls keep files open; move the loop body into a function.
- Arguments are evaluated immediately; wrap in `func() { ... }()` to read the latest values.
:::
