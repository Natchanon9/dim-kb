---
title: "fs Package"
description: "The fs.FS interface for read-only file trees"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`io/fs` defines `fs.FS`, a read-only filesystem interface. Code that accepts an `fs.FS` works with the real disk (`os.DirFS`), embedded files (`embed.FS`), zip files and in-memory test data.

## Example

```go title="fsys.go"
var fsys fs.FS = fstest.MapFS{ // in-memory, handy in tests
	"a.txt":     {Data: []byte("A")},
	"dir/b.txt": {Data: []byte("B")},
}
fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		fmt.Println(p) // a.txt, dir/b.txt
	}
	return nil
})
data, _ := fs.ReadFile(fsys, "dir/b.txt") // B
```

| Source | Gives an `fs.FS` |
| --- | --- |
| Disk | `os.DirFS("templates")` |
| Binary | `embed.FS` |
| Tests | `fstest.MapFS` |

:::caution[Pitfalls]
- `fs.FS` paths always use `/` and are relative: no leading `/`, no `..`.
:::
