---
title: "ioutil Alternatives"
description: "Where every io/ioutil function moved in Go 1.16"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`io/ioutil` is deprecated since Go 1.16. Its functions moved to `io` and `os` with the same behaviour.

## Mapping

| Old (`ioutil`) | New |
| --- | --- |
| `ReadAll` | `io.ReadAll` |
| `ReadFile` / `WriteFile` | `os.ReadFile` / `os.WriteFile` |
| `ReadDir` | `os.ReadDir` (returns `[]os.DirEntry`, cheaper) |
| `TempDir` / `TempFile` | `os.MkdirTemp` / `os.CreateTemp` |
| `NopCloser` | `io.NopCloser` |
| `Discard` | `io.Discard` |

```go title="readdir.go"
dir, _ := os.MkdirTemp("", "list")
defer os.RemoveAll(dir)
entries, _ := os.ReadDir(dir) // sorted by name
for _, e := range entries {
	fmt.Println(e.Name())
}
```

:::caution[Pitfalls]
- `os.ReadDir` returns `DirEntry`, not `FileInfo`; call `e.Info()` if you need size or mod time.
:::
