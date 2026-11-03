---
title: "Code Coverage"
description: "See which lines your tests run"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

Coverage shows which statements the tests executed. Use it to find untested branches, not as a score to maximise.

## Commands

```sh
go test -cover ./...                          # percentage per package
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out                 # open an annotated view
go tool cover -func=cover.out                 # per-function table
go test -coverpkg=./... -coverprofile=c.out ./... # count cross-package calls
```

## Example

```go title="classify_test.go"
func TestClassify(t *testing.T) {
	for in, want := range map[int]string{-1: "negative", 0: "zero", 5: "positive"} {
		if got := Classify(in); got != want {
			t.Errorf("Classify(%d) = %q, want %q", in, got, want)
		}
	}
}
// coverage: 100.0% of statements
```

Since Go 1.20, `go build -cover` also instruments whole binaries for integration-test coverage.

:::caution[Pitfalls]
- 100% coverage doesn't mean correct; a line can run without its result being checked.
:::
