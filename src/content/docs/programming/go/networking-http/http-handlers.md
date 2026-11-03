---
title: "HTTP Handlers"
description: "Handler, HandlerFunc and handlers with dependencies"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A handler gets a `ResponseWriter` to write the response and a `*Request` to read input. Any type with `ServeHTTP(w, r)` is an `http.Handler`.

## A struct handler with dependencies

```go title="greeter.go"
type greeter struct{ greeting string } // DB, logger, config go here

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s", g.greeting, r.URL.Query().Get("name"))
}
// GET /?name=ann → hi, ann
```

## Reading the request

| Need | Use |
| --- | --- |
| Path param | `r.PathValue("id")` (Go 1.22+) |
| Query string | `r.URL.Query().Get("q")` |
| Header | `r.Header.Get("Authorization")` |
| JSON body | `json.NewDecoder(r.Body).Decode(&v)` |
| Form | `r.FormValue("email")` |
| Cancellation | `r.Context()` |

## Writing the response

Set headers, then `w.WriteHeader(status)`, then write the body. Writing the body first sends `200` implicitly.

:::caution[Pitfalls]
- Setting headers after `WriteHeader` or the first `Write` has no effect.
- After `http.Error(...)`, `return`; otherwise the handler keeps writing.
:::
