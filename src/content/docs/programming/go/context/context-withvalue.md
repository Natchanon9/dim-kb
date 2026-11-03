---
title: "context.WithValue"
description: "Attach request-scoped values like request IDs"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`WithValue` attaches a key-value pair to a context. Use it only for request-scoped data that crosses API boundaries: request ID, trace span, authenticated user.

## Example

```go title="value.go"
type requestIDKey struct{} // unexported key type: no collisions

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

func requestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey{}).(string)
	return id, ok
}

ctx := withRequestID(context.Background(), "req-42")
requestID(ctx)                  // req-42 true
requestID(context.Background()) // "" false
```

Wrap access in typed helper functions so callers never touch keys or type assertions.

:::caution[Pitfalls]
- Don't pass function parameters through context (DB handles, config, optional args); make them explicit.
- Never use a plain `string` as the key; another package can collide with it.
:::
