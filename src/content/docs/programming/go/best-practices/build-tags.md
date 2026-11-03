---
title: "Build Tags"
description: "Include files only for certain OSes, versions or custom tags"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A `//go:build` line at the top of a file decides whether the file is part of the build.

## Example: two editions

```go title="edition_free.go"
//go:build !premium

package app

func Edition() string { return "free" }
```

```go title="edition_premium.go"
//go:build premium

package app

func Edition() string { return "premium" }
```

```sh
go build                 # Edition() → free
go build -tags premium   # Edition() → premium
```

## Common constraints

| Constraint | Meaning |
| --- | --- |
| `//go:build linux` | Only on Linux (`windows`, `darwin` …) |
| `//go:build amd64 \|\| arm64` | Architectures |
| `//go:build go1.23` | Go version ≥ 1.23 |
| `//go:build integration` | Custom tag, e.g. slow tests |
| `//go:build cgo` | Only when cgo is enabled |

File names work too: `file_linux.go`, `file_test.go`.

:::caution[Pitfalls]
- The line must come before `package` and be followed by a blank line; otherwise it's just a comment.
:::
