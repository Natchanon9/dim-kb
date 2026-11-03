---
title: "Wrapping Errors"
description: "Add context with %w and keep the original error inside"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Wrapping adds context to an error while keeping the original inside it, so callers can still detect what really happened.

## %w vs %v

```go title="wrap.go"
base := errors.New("connection refused")
wrapped := fmt.Errorf("query users: %w", base) // %w keeps base inside
flat := fmt.Errorf("query users: %v", base)    // %v only copies the text
fmt.Println(wrapped)                            // query users: connection refused
errors.Unwrap(wrapped) == base // true
errors.Unwrap(flat) == nil     // true
```

Wrapped errors form a chain; `errors.Is` and `errors.As` walk it.

## When not to wrap

Use `%v` when you don't want callers to depend on an internal error (for example, hiding which database driver you use behind a package boundary).

:::caution[Pitfalls]
- `fmt.Errorf` can take several `%w` (Go 1.20+), creating a tree rather than a chain.
:::
