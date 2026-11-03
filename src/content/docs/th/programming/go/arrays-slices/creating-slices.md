---
title: "การสร้าง Slice"
description: "literal, make พร้อม length และ capacity, nil slice, slice จาก array"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

เลือกวิธีสร้างตามสิ่งที่รู้ล่วงหน้า: รู้ค่า รู้ขนาด หรือยังไม่รู้อะไรเลย

## รูปแบบ

```go title="create.go"
a := []int{1, 2}        // literal
b := make([]int, 3)     // len 3 เป็นศูนย์หมด
c := make([]int, 0, 10) // len 0 มีที่ว่าง 10
var d []int             // nil slice
fmt.Println(a, b, len(c), cap(c), d == nil)
// [1 2] [0 0 0] 0 10 true
```

| รู้อะไร | ใช้ |
| --- | --- |
| รู้ค่า | literal `[]T{...}` |
| รู้ความยาวแน่นอน | `make([]T, n)` แล้วกำหนด `s[i]` |
| รู้คร่าวๆ ว่าจะ append กี่ตัว | `make([]T, 0, n)` แล้ว `append` |
| ยังไม่รู้อะไร | `var s []T` |

:::caution[ข้อควรระวัง]
- `make([]T, n)` แล้ว `append` จะต่อ **หลัง** ศูนย์ n ตัว ถ้าจะ append ให้ใช้ `make([]T, 0, n)`
:::
