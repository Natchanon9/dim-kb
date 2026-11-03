---
title: "Creating Errors"
description: "errors.New for fixed messages, fmt.Errorf for formatted ones"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Two functions cover most error creation.

## Example

```go title="create.go"
e1 := errors.New("not found")            // fixed message
e2 := fmt.Errorf("user %d not found", 7) // formatted message
fmt.Println(e1) // not found
fmt.Println(e2) // user 7 not found
```

| Need | Use |
| --- | --- |
| Fixed message, maybe a sentinel | `errors.New` |
| Include values | `fmt.Errorf("... %d", v)` |
| Add context to another error | `fmt.Errorf("...: %w", err)` |
| Extra fields callers inspect | A custom error type |

## Message style

Lowercase, no trailing punctuation, no "error:" prefix, because messages get chained: `load config: open file: permission denied`.

:::caution[Pitfalls]
- Two `errors.New("x")` calls make two **different** errors; `==` between them is false.
:::
