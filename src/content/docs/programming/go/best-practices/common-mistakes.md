---
title: "Common Mistakes"
description: "Bugs that show up again and again in Go code"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Most of these have their own page; this is the checklist to scan during review.

## Two classics

```go title="mistakes.go"
m := map[string]int{} // not `var m map[string]int`: writing to a nil map panics
m["a"]++

func find(ok bool) error {
	if !ok {
		return &notFound{}
	}
	return nil // a literal nil, never a typed nil pointer
}
```

## Checklist

| Mistake | Fix |
| --- | --- |
| Writing to a nil map | `make` it first |
| Returning a typed nil as `error` | `return nil` |
| `defer` inside a long loop | Move the body into a function |
| Ignoring errors with `_` | Handle or explicitly justify |
| Goroutine with no exit | Context or done channel |
| `http.Get` / client without timeout | `&http.Client{Timeout: ...}` |
| Not closing `resp.Body` / `rows` | `defer x.Close()` |
| Modifying a `range` copy | Index into the slice |
| Sharing a backing array after `append` | Copy when ownership is unclear |
| Data races | `go test -race` |
| Comparing errors with `==` | `errors.Is` / `errors.As` |

:::caution[Pitfalls]
- Most of these are caught by `go vet`, staticcheck and `-race`; run them in CI.
:::
