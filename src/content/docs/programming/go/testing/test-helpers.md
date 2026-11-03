---
title: "Test Helpers"
description: "t.Helper, t.Cleanup, t.TempDir and t.Setenv"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Helpers remove repetition from tests. `testing.T` also ships helpers for common setup and teardown.

## Example

```go title="helpers_test.go"
func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper() // failures point at the caller's line
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWithHelpers(t *testing.T) {
	t.Setenv("APP_ENV", "test") // restored afterwards
	dir := t.TempDir()          // deleted afterwards
	t.Cleanup(func() { /* close DBs, stop servers */ })

	assertEqual(t, Greet("Ann"), "Hello, Ann!")
}
```

| Helper | Does |
| --- | --- |
| `t.Helper()` | Hide this function from failure line numbers |
| `t.Cleanup(f)` | Run f when the test ends (LIFO) |
| `t.TempDir()` | Fresh directory, auto-removed |
| `t.Setenv(k, v)` | Set env var, restore after |

For richer assertions many teams use `github.com/stretchr/testify` (`require`, `assert`).

:::caution[Pitfalls]
- `t.Setenv` can't be used with `t.Parallel()`.
:::
