---
title: "go generate"
description: "Run code generators declared in source comments"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`go generate` scans for `//go:generate` comments and runs the commands. Generated files are committed, so normal builds don't need the tools.

## Example: stringer

```go title="color.go"
//go:generate stringer -type=Color

type Color int

const (
	Red Color = iota
	Green
	Blue
)
```

```sh
go install golang.org/x/tools/cmd/stringer@latest
go generate ./...   # writes color_string.go with func (Color) String()
```

```go
fmt.Println(Red, Blue, Color(7)) // Red Blue Color(7)
```

## Common generators

| Tool | Generates |
| --- | --- |
| `stringer` | `String()` for enums |
| `mockgen` | Mocks from interfaces |
| `sqlc` | Typed Go from SQL |
| `protoc-gen-go` | Protocol Buffers / gRPC |
| `swag` | Swagger docs from comments |

:::caution[Pitfalls]
- `go build` never runs `go generate`; rerun it after changing the source and check CI that generated files are up to date.
:::
