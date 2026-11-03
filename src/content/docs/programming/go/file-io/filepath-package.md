---
title: "filepath Package"
description: "Build and split OS file paths portably"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`path/filepath` handles file paths with the right separator for the OS (`/` or `\`). Never build paths with string concatenation.

## Common functions

```go title="paths.go"
p := filepath.Join("data", "users", "..", "orders", "2024.csv") // also cleans ".."
filepath.Base(p) // 2024.csv
filepath.Ext(p)  // .csv
filepath.Dir(p)  // data/orders
filepath.Match("*.csv", "report.csv") // true
```

| Function | Use |
| --- | --- |
| `Join`, `Clean` | Build a normalised path |
| `Base`, `Dir`, `Ext` | Split it |
| `Abs`, `Rel` | Absolute / relative paths |
| `Glob`, `Match` | Wildcards |
| `WalkDir` | Visit every file in a tree |

`path` (without `filepath`) is for slash-separated paths like URLs and `fs.FS` paths.

:::caution[Pitfalls]
- Joining user input into a path can escape your directory (`../../etc/passwd`). Use `filepath.IsLocal` (Go 1.20+) or `os.Root` (Go 1.24+) to stay inside a base directory.
:::
