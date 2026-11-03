---
title: "พื้นฐาน Slice"
description: "มุมมองแบบยืดได้บน array: literal, index, slice expression, len"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

slice คือมุมมองที่ขยายได้บน array เป็น type แบบ list ที่ใช้ทุกที่ใน Go

## พื้นฐาน

```go title="slice.go"
s := []int{1, 2, 3} // ไม่มีความยาวในวงเล็บ → slice
s = append(s, 4)
fmt.Println(s, len(s), s[1:3]) // [1 2 3 4] 4 [2 3]
```

- `s[i]` อ่าน/เขียน element ตัวเดียว
- `s[low:high]` สร้าง slice ใหม่ตั้งแต่ `low` ถึงก่อน `high` ละปลายใดปลายหนึ่งได้: `s[:2]`, `s[1:]`
- `len(s)` คือจำนวน element ส่วน `cap(s)` คือที่ว่างก่อนต้องจองใหม่

## Slice vs array

| | Array `[3]int` | Slice `[]int` |
| --- | --- | --- |
| ความยาว | คงที่ เป็นส่วนหนึ่งของ type | ขยายด้วย `append` |
| กำหนดค่า | copy element | copy header แต่แชร์ข้อมูล |
| เทียบด้วย `==` | ได้ | เทียบได้แค่กับ `nil` |

:::caution[ข้อควรระวัง]
- ใช้ผลลัพธ์ของ `append` เสมอ: `s = append(s, x)`
- sub-slice แชร์ memory กับตัวต้นทาง เขียนผ่านตัวหนึ่ง อีกตัวก็เห็น
:::
