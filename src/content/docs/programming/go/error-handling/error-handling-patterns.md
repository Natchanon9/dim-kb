---
title: "Error Handling Patterns"
description: "Check immediately, return early, add context once per layer"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

The Go pattern is: check the error right after the call, add context, return early. The happy path stays at the left margin.

## Example

```go title="config.go"
func loadConfig(path string) (string, error) {
	data, err := readFile(path)
	if err != nil {
		return "", fmt.Errorf("load config: %w", err)
	}
	return data, nil
}
// loadConfig("") → load config: empty path
```

## Rules

- **Check every error.** Ignore one only on purpose, with `_` and a comment.
- **Add context on the way up:** what this layer was doing, not what failed below.
- **Handle once:** either log it or return it, not both, or the same error is logged at every layer.
- **Decide at the top:** HTTP handlers or `main` turn errors into status codes, messages or exits.

:::caution[Pitfalls]
- `if err != nil { return err }` everywhere without context makes errors like `EOF` impossible to trace.
:::
