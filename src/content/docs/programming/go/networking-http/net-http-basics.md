---
title: "net/http Basics"
description: "A production-ready HTTP server and client in the standard library"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`net/http` contains a full HTTP/1.1 and HTTP/2 server and client. Many Go services use it directly with no framework.

## Minimal server

```go title="main.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
})
log.Fatal(http.ListenAndServe(":8080", mux))
// GET /hello → 200 hello
```

## The pieces

| Piece | Role |
| --- | --- |
| `http.Handler` | Anything with `ServeHTTP(w, r)` |
| `http.HandlerFunc` | Turns a function into a Handler |
| `http.ServeMux` | Router: pattern → handler |
| `http.Server` | Listener with timeouts and shutdown |
| `http.Client` | Makes outgoing requests |

Each request runs in its own goroutine, so handlers must be safe for concurrent use.

:::caution[Pitfalls]
- `http.ListenAndServe` has no timeouts; in production build an `http.Server` with `ReadHeaderTimeout` (and usually `ReadTimeout`, `WriteTimeout`, `IdleTimeout`).
:::
