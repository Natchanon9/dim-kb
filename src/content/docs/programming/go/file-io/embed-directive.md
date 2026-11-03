---
title: "embed Directive"
description: "Bake files into the binary with //go:embed"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`//go:embed` (Go 1.16+) includes files in the compiled binary, so templates, SQL migrations and static assets ship with the single executable.

## Example

```go title="assets.go"
import "embed"

//go:embed static/hello.txt
var hello string // file contents as a string ([]byte works too)

//go:embed static
var staticFS embed.FS // a whole directory as an fs.FS

fmt.Print(hello) // hello from embed
entries, _ := fs.ReadDir(staticFS, "static")
```

Serve embedded assets directly: `http.FileServerFS(staticFS)` (Go 1.22+).

:::caution[Pitfalls]
- The directive must sit right above a package-level `var`, and the package must import `embed`.
- Paths are relative to the source file and can't use `..`; files starting with `.` or `_` are skipped unless you use `all:`.
:::
