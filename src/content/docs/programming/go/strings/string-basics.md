---
title: "String Basics"
description: "Immutable UTF-8 bytes, len, indexing, raw strings and concatenation"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A Go string is an immutable sequence of bytes, usually UTF-8 text. Operations that "change" a string create a new one.

## Basics

```go title="string.go"
s := "go"
t := s + "lang"    // creates a new string
raw := `C:\path\n` // raw string: no escapes, can span lines
fmt.Println(t, len("héllo"), s[0], raw)
// golang 6 103 C:\path\n
```

- `len(s)` counts **bytes**, not characters: `é` takes 2 bytes.
- `s[i]` is a **byte** (`uint8`), not a character.
- `s[i:j]` slices bytes and shares memory with `s`.
- Strings compare with `==` and `<` byte by byte.

:::caution[Pitfalls]
- `s[0] = 'G'` doesn't compile; convert to `[]byte` or `[]rune`, change, convert back.
- Slicing by byte index can cut a multi-byte character in half (Thai is 3 bytes per character).
:::
