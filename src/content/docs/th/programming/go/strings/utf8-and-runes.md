---
title: "UTF-8 และ Rune"
description: "byte vs ตัวอักษร, rune และการนับตัวอักษรให้ถูก"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`rune` คือ Unicode code point (`int32`) UTF-8 เก็บแต่ละ rune ใน 1–4 byte ข้อความที่ไม่ใช่ ASCII จึงมีความยาว byte ไม่เท่าจำนวนตัวอักษร

## Byte vs rune

```go title="runes.go"
s := "สวัสดี"
r := []rune(s)
fmt.Println(len(s), utf8.RuneCountInString(s), r[0] == 'ส', string(r[:2]))
// 18 6 true สว
```

| ต้องการ | ใช้ |
| --- | --- |
| ความยาว byte (เก็บข้อมูล, network) | `len(s)` |
| จำนวนตัวอักษร | `utf8.RuneCountInString(s)` |
| index ตามตัวอักษร | `[]rune(s)[i]` |
| ตรวจ input | `utf8.ValidString(s)` |

rune literal ใช้ single quote: `'ก'` ส่วน byte ที่ไม่ถูกต้องจะถอดเป็น `utf8.RuneError` (`�`)

:::caution[ข้อควรระวัง]
- สระและวรรณยุกต์ไทยที่ซ้อนบนล่าง (`ั`, `ี`, `่`) เป็น rune แยก "ตัวที่ตาเห็น" จึงอาจน้อยกว่าจำนวน rune
- `[]rune(s)` จอง memory ใหม่ ถ้าอยู่ในส่วนที่ทำงานหนักให้ใช้ `range` แทน
:::
