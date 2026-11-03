---
title: "Basic Types"
description: "integer, float, bool, string, byte และ rune"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go มี built-in type ไม่กี่ตัว ใช้ `int`, `float64`, `string` และ `bool` เป็นหลัก เว้นแต่มีเหตุผลต้องใช้ตัวอื่น

## Type ที่มี

| ประเภท | Type | หมายเหตุ |
| --- | --- | --- |
| Integer | `int`, `int8`…`int64`, `uint`…`uint64` | `int` เป็น 64-bit บนเครื่อง 64-bit |
| Float | `float32`, `float64` | ค่าเริ่มต้นคือ `float64` |
| Complex | `complex64`, `complex128` | แทบไม่ได้ใช้ |
| ข้อความ | `string` | byte แบบ UTF-8 ที่แก้ไขไม่ได้ |
| Alias | `byte` = `uint8`, `rune` = `int32` | rune คือ Unicode code point 1 ตัว |
| Boolean | `bool` | มีแค่ `true` / `false` ใช้ตัวเลขแทนไม่ได้ |

```go title="types.go"
var i int = 42
var f float64 = 3.14
var b byte = 'A'
var r rune = 'ก'
var s string = "hi"
fmt.Printf("%T %T %T %T %T\n", i, f, b, r, s)
// int float64 uint8 int32 string
```

:::caution[ข้อควรระวัง]
- ผสม type ต้องแปลงเอง: `int + int64` compile ไม่ผ่าน
- float ไม่แม่นยำ อย่าเก็บเงินเป็น `float64` ให้เก็บเป็นสตางค์แบบ integer
:::
