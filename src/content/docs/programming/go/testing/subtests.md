---
title: "Subtests"
description: "Named sub-cases with t.Run, filtering and parallelism"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`t.Run(name, func(t *testing.T) {...})` creates a subtest with its own name, pass/fail state and `Fatal` scope.

## Example

```go title="sub_test.go"
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		t.Parallel() // run cases concurrently
		if got := Abs(tt.in); got != tt.want {
			t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
		}
	})
}
```

```sh
go test -run 'TestAbsSubtests/negative' -v
```

## Benefits

- Run or debug one case by name.
- `t.Fatal` stops only that subtest.
- Shared setup before the loop, per-case setup inside.

:::caution[Pitfalls]
- With `t.Parallel()` on Go < 1.22, capture `tt := tt` first; since 1.22 the loop variable is per iteration.
:::
