---
title: "Common Interfaces"
description: "fmt.Stringer, error, io.Reader, io.Writer and friends"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

A few standard interfaces appear everywhere. Implementing them lets your types plug into the standard library.

## Stringer and io

```go title="common.go"
type Temp float64

func (t Temp) String() string { return fmt.Sprintf("%.1f°C", float64(t)) }

fmt.Println(Temp(21.5)) // 21.5°C

r := strings.NewReader("hello") // io.Reader
var sb strings.Builder          // io.Writer
n, _ := io.Copy(&sb, r)
fmt.Println(n, sb.String()) // 5 hello
```

| Interface | Method | Used by |
| --- | --- | --- |
| `fmt.Stringer` | `String() string` | `fmt` printing |
| `error` | `Error() string` | All error handling |
| `io.Reader` | `Read(p []byte) (n int, err error)` | Files, HTTP bodies, gzip, JSON decoder |
| `io.Writer` | `Write(p []byte) (n int, err error)` | Files, responses, buffers, hashes |
| `io.Closer` | `Close() error` | Anything holding a resource |
| `sort.Interface` | `Len`, `Less`, `Swap` | `sort.Sort` (mostly replaced by `slices.SortFunc`) |

:::caution[Pitfalls]
- `Read` can return data **and** `io.EOF` in the same call; process `n` bytes before checking the error.
:::
