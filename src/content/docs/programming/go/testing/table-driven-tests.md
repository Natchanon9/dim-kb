---
title: "Table-Driven Tests"
description: "One loop over a slice of cases"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Put inputs and expected outputs in a slice of structs and loop over it. Adding a case is one line.

## Example

```go title="abs_test.go"
func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		in, want int
	}{
		{"positive", 3, 3},
		{"negative", -3, 3},
		{"zero", 0, 0},
	}
	for _, tt := range tests {
		if got := Abs(tt.in); got != tt.want {
			t.Errorf("%s: Abs(%d) = %d, want %d", tt.name, tt.in, got, tt.want)
		}
	}
}
```

Combine with `t.Run` (next page) so each case is a named subtest you can run on its own.

:::caution[Pitfalls]
- Include a `name` field; "case 3 failed" is hard to find.
- Use `t.Errorf`, not `t.Fatalf`, inside the loop so one failure doesn't hide the others.
:::
