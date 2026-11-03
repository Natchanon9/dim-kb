---
title: "Request Context"
description: "r.Context() for cancellation and request-scoped values"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Every request carries a context. It is cancelled when the client disconnects, the HTTP/2 stream is reset, or the server shuts down.

## Example

```go title="context.go"
type ridKey struct{}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ridKey{}, "req-7")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func report(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(time.Second): // slow work
		fmt.Fprint(w, "done")
	case <-r.Context().Done():
		log.Println(r.Context().Value(ridKey{}), "stopped:", r.Context().Err())
	}
}
// client disconnects → req-7 stopped: context canceled
```

Pass `r.Context()` to every DB query and outgoing call so abandoned requests stop using resources.

:::caution[Pitfalls]
- Don't use `r.Context()` for work that must finish after you respond; it is cancelled when the handler returns.
:::
