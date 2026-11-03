---
title: "bufio Package"
description: "Buffered reading line by line and buffered writing"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`bufio` wraps a reader or writer with a buffer, so many small reads or writes become a few big system calls.

## Reading lines

```go title="scan.go"
sc := bufio.NewScanner(strings.NewReader("a\nb\nc")) // or an *os.File
n, last := 0, ""
for sc.Scan() {
	n++
	last = sc.Text()
}
fmt.Println(n, last, sc.Err()) // 3 c <nil>
```

`sc.Split(bufio.ScanWords)` splits by words instead of lines.

## Buffered writing

```go title="write.go"
var sb strings.Builder
w := bufio.NewWriter(&sb)
w.WriteString("buffered")
sb.Len() // 0: still in the buffer
w.Flush()
sb.String() // buffered
```

:::caution[Pitfalls]
- `Scanner` rejects lines over 64 KB by default; raise it with `sc.Buffer(buf, max)`.
- Forgetting `Flush()` loses the last chunk of output.
:::
