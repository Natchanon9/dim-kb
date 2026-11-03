---
title: "errors.Is & errors.As"
description: "Check for a specific error value or type anywhere in the chain"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`errors.Is` asks "is this specific error anywhere in the chain?". `errors.As` asks "is there an error of this type, and give it to me".

## Example

```go title="isas.go"
var ErrNotFound = errors.New("not found")

err := fmt.Errorf("get user: %w", ErrNotFound)
errors.Is(err, ErrNotFound) // true

err = fmt.Errorf("save: %w", validateEmail("bob"))
var ve *ValidationError
if errors.As(err, &ve) {
	fmt.Println(ve.Field) // email
}
```

| Function | Compares | Typical target |
| --- | --- | --- |
| `errors.Is(err, target)` | Values | Sentinels: `io.EOF`, `sql.ErrNoRows` |
| `errors.As(err, &target)` | Types | Custom types: `*ValidationError`, `*fs.PathError` |

:::caution[Pitfalls]
- `errors.As` needs a **pointer** to a variable of the target type: `&ve`, where `ve` is `*ValidationError`.
- `err == ErrNotFound` fails once the error is wrapped; always use `errors.Is`.
:::
