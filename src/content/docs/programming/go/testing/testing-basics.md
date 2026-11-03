---
title: "Testing Basics"
description: "Write and run tests with go test"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Tests live in `_test.go` files next to the code. A test is a function `TestXxx(t *testing.T)`; `go test` finds and runs them. No framework needed.

## Example

```go title="add_test.go"
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", got)
	}
}
```

```sh
go test ./...           # all packages
go test -run TestAdd -v # one test, verbose
go test -race ./...     # with the race detector
go test -count=1 ./...  # skip the test cache
```

## Reporting

| Call | Effect |
| --- | --- |
| `t.Errorf` / `t.Error` | Mark failed, keep going |
| `t.Fatalf` / `t.Fatal` | Mark failed, stop this test |
| `t.Skip` | Skip (e.g. needs a database) |
| `t.Log` | Print only on failure or with `-v` |

Message convention: `Func(args) = got, want want`.

:::caution[Pitfalls]
- `t.Fatal` from a goroutine you started doesn't stop the test; report back to the test goroutine instead.
:::
