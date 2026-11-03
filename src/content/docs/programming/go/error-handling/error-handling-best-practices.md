---
title: "Error Handling Best Practices"
description: "A checklist for errors in production code"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Good error handling makes failures easy to trace and easy to react to.

## Checklist

- **Wrap with context** at each layer: `fmt.Errorf("charge order %d: %w", id, err)`.
- **Handle once:** log **or** return, never both.
- **Compare with `errors.Is`/`errors.As`,** never by message.
- **Sentinels for "which", types for "details".**
- **Lowercase messages,** no trailing punctuation.
- **Don't leak internals** to API clients: map errors to safe messages and status codes at the edge.
- **Report several problems** with `errors.Join` (Go 1.20+):

```go title="join.go"
err := errors.Join(
	errors.New("name required"),
	errors.New("email invalid"),
)
fmt.Println(err)
// name required
// email invalid
```

`errors.Is` and `errors.As` search every joined error.

:::caution[Pitfalls]
- Panicking for expected failures and ignoring errors with `_` are the two most common review comments on Go code.
:::
