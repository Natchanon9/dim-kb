---
title: "HTTP Middleware"
description: "Wrap handlers to add logging, auth and recovery"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Middleware is a function that takes a handler and returns a new one. It runs code before and after the wrapped handler, or stops the request early.

## Example

```go title="middleware.go"
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("->", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func requireKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != "secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return // stop the chain
		}
		next.ServeHTTP(w, r)
	})
}

h := logging(requireKey(dataHandler)) // outermost runs first
// GET /data without key → "-> GET /data", 401
```

Common middleware: logging, request ID, auth, CORS, panic recovery, rate limiting, timeouts (`http.TimeoutHandler`).

:::caution[Pitfalls]
- To log the status code, wrap `ResponseWriter` in a struct that records `WriteHeader`.
:::
