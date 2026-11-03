---
title: "String Iteration"
description: "range over runes vs indexing bytes"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`for i, r := range s` decodes UTF-8 and gives runes. A classic `for i := 0; i < len(s); i++` gives raw bytes.

## Both ways

```go title="iterate.go"
s := "aก"
for i, r := range s { // by rune; i is the byte offset
	fmt.Println(i, string(r), r) // 0 a 97, 1 ก 3585
}
for i := 0; i < len(s); i++ { // by byte
	fmt.Print(s[i], " ") // 97 224 184 129
}
```

Use `range` for text. Use byte loops only for ASCII protocols or binary parsing.

:::caution[Pitfalls]
- The index from `range` jumps (0, 1, 4 …) because it is a byte offset, not a character position.
:::
