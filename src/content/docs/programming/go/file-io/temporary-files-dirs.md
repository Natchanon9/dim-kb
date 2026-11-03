---
title: "Temporary Files & Dirs"
description: "Create scratch files safely and clean them up"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`os.CreateTemp` and `os.MkdirTemp` create uniquely named files and directories under the system temp folder. You are responsible for removing them.

## Example

```go title="temp.go"
f, err := os.CreateTemp("", "upload-*.txt") // * → random string
if err != nil {
	return err
}
defer os.Remove(f.Name())
f.WriteString("tmp")
f.Close()
// name looks like /tmp/upload-123456789.txt
```

## In tests

```go
dir := t.TempDir() // removed automatically when the test ends
```

:::caution[Pitfalls]
- Always `defer os.Remove`/`os.RemoveAll`; temp folders on servers fill up silently.
- Don't build temp names yourself with fixed names; concurrent runs collide.
:::
