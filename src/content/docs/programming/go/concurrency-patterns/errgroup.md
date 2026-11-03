---
title: "errgroup"
description: "Run goroutines, collect the first error and cancel the rest"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`golang.org/x/sync/errgroup` is a WaitGroup that returns an error. With `WithContext`, the first failure cancels a shared context so the other goroutines can stop.

## Example

```go title="errgroup.go"
g, ctx := errgroup.WithContext(context.Background())
g.SetLimit(2) // at most 2 at once
for _, u := range []string{"a", "bad", "c"} {
	g.Go(func() error { return fetch(ctx, u) })
}
if err := g.Wait(); err != nil {
	fmt.Println(err) // fetch bad: failed
}
```

`fetch` should watch `ctx.Done()` so it stops once another call failed.

## When to use

Parallel calls where any failure fails the whole operation: loading data from several services for one response, uploading parts of a file.

:::caution[Pitfalls]
- `Wait` returns only the **first** error; use `errors.Join` yourself if you need all of them.
- Use the derived `ctx` inside `g.Go`, not the parent.
:::
