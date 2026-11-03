---
title: "append และ copy"
description: "ขยาย slice ด้วย append และทำสำเนาข้อมูลด้วย copy"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`append` เพิ่ม element และอาจจอง backing array ใหม่ ส่วน `copy` ย้าย element ระหว่าง slice ที่มีอยู่แล้ว

## append

```go title="append.go"
s := make([]int, 0, 2)
s = append(s, 1, 2)
before := cap(s)
s = append(s, 3) // เกิน capacity: ได้ backing array ใหม่ที่ใหญ่กว่า
fmt.Println(s, before, cap(s) > before) // [1 2 3] 2 true
```

ต่อ slice เข้าด้วยกันด้วย `append(a, b...)`

## copy

```go title="copy.go"
dst := make([]int, 2)
n := copy(dst, []int{7, 8, 9}) // copy เท่ากับ min(len(dst), len(src))
fmt.Println(dst, n)            // [7 8] 2
```

ถ้าจะทำสำเนาทั้ง slice ใช้ `slices.Clone(s)` สั้นที่สุด

:::caution[ข้อควรระวัง]
- `copy` ลง slice ที่ยาว 0 จะไม่ copy อะไรเลย ให้กำหนดขนาด `dst` ด้วย `make([]T, len(src))` ก่อน
- append ใน loop ที่ทำงานหนักโดยไม่จองขนาดไว้ก่อนจะจอง memory ซ้ำหลายรอบ ใช้ `make([]T, 0, n)`
:::
