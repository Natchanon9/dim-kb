---
title: "Pointer กับ Function"
description: "ให้ function แก้ตัวแปรของผู้เรียก และคืน pointer อย่างปลอดภัย"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

argument เป็นสำเนา ส่ง pointer ไปแล้ว function จะแก้ตัวแปรของผู้เรียกได้

## ตัวอย่าง

```go title="funcs.go"
func reset(n *int) { *n = 0 }

func newInt(v int) *int { return &v } // ใน Go ทำได้ปลอดภัย

x := 5
reset(&x)
p := newInt(3)
fmt.Println(x, *p) // 0 3
```

## คืน pointer ของตัวแปร local

ต่างจากภาษา C การคืน `&v` ของตัวแปร local ทำได้ escape analysis ของ compiler จะย้าย `v` ไปไว้บน heap ถ้ามันต้องอยู่นานกว่า function

:::caution[ข้อควรระวัง]
- pointer ไม่ได้ทำให้โค้ดเร็วขึ้นเสมอ struct เล็กๆ ที่ส่งเป็นค่ามักไม่ต้องจองบน heap และเร็วกว่า
- function ที่รับ pointer ควรบอกให้ชัดว่าจะเก็บ pointer นั้นไว้ใช้ต่อหลัง return หรือไม่
:::
