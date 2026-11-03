---
title: "UTF-8 & Runes"
description: "Bytes vs characters, rune, and counting characters correctly"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A `rune` is a Unicode code point (`int32`). UTF-8 stores each rune in 1–4 bytes, so byte length and character count differ for non-ASCII text.

## Bytes vs runes

```go title="runes.go"
s := "สวัสดี"
r := []rune(s)
fmt.Println(len(s), utf8.RuneCountInString(s), r[0] == 'ส', string(r[:2]))
// 18 6 true สว
```

| Need | Use |
| --- | --- |
| Byte length (storage, network) | `len(s)` |
| Character count | `utf8.RuneCountInString(s)` |
| Index by character | `[]rune(s)[i]` |
| Validate input | `utf8.ValidString(s)` |

A rune literal uses single quotes: `'ก'`. Invalid bytes decode as `utf8.RuneError` (`�`).

:::caution[Pitfalls]
- Thai combining vowels and tone marks (`ั`, `ี`, `่`) are separate runes, so "characters you see" can be fewer than runes.
- `[]rune(s)` allocates; avoid it in hot paths and use `range` instead.
:::
