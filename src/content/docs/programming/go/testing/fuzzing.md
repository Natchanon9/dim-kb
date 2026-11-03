---
title: "Fuzzing"
description: "Let go test generate inputs to find crashes (Go 1.18+)"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A fuzz test `FuzzXxx(f *testing.F)` starts from seed inputs and mutates them to find inputs that break a property you assert.

## Example

```go title="reverse_test.go"
func FuzzReverse(f *testing.F) {
	f.Add("hello") // seed corpus
	f.Add("สวัสดี")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}
		if got := Reverse(Reverse(s)); got != s {
			t.Errorf("Reverse twice = %q, want %q", got, s)
		}
	})
}
```

```sh
go test ./...                          # runs seeds only
go test -fuzz=FuzzReverse -fuzztime=30s # actually fuzzes
```

Failing inputs are saved to `testdata/fuzz/FuzzReverse/` and replayed by every later `go test`.

## Good targets

Parsers, decoders, validators, anything taking untrusted input. Assert properties: round-trips, no panics, invariants.

:::caution[Pitfalls]
- Fuzzing runs one target per command; CI usually just runs the seeds.
:::
