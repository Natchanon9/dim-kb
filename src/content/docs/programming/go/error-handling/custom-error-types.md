---
title: "Custom Error Types"
description: "Error structs that carry extra fields"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

When callers need details (which field, which status code), return a struct that implements `error`.

## Example

```go title="validation.go"
type ValidationError struct {
	Field, Reason string
}

func (e *ValidationError) Error() string { return e.Field + ": " + e.Reason }

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Reason: "missing @"}
	}
	return nil
}
// validateEmail("bob") → email: missing @
```

Callers extract it with `errors.As`. If your type wraps another error, add `Unwrap() error` so `Is`/`As` can see inside.

:::caution[Pitfalls]
- Return `nil`, not a nil `*ValidationError`, on success. A typed nil makes `err != nil` true.
- Use a pointer receiver and return `&T{}`, so `errors.As` targets are consistent.
:::
