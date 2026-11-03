---
title: "os Package"
description: "Read, write, open and inspect files"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`os` is the low-level door to the filesystem: whole-file helpers for small files, and `*os.File` for streaming.

## Whole-file helpers

```go title="files.go"
path := filepath.Join(dir, "note.txt")
err := os.WriteFile(path, []byte("hello\n"), 0o644)
data, err := os.ReadFile(path) // "hello\n"

_, err = os.Stat(filepath.Join(dir, "missing"))
errors.Is(err, os.ErrNotExist) // true
```

## Opening files

| Call | Use |
| --- | --- |
| `os.Open(p)` | Read only |
| `os.Create(p)` | Create or truncate, write |
| `os.OpenFile(p, os.O_APPEND\|os.O_CREATE\|os.O_WRONLY, 0o644)` | Append, custom flags |
| `os.MkdirAll(p, 0o755)` | Create directories |
| `os.Remove` / `os.RemoveAll` | Delete |

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close()
```

:::caution[Pitfalls]
- Check the error from `Close` on files you **wrote**; it can report a failed flush.
- `ReadFile` loads everything into memory; stream big files with `bufio`/`io.Copy`.
:::
