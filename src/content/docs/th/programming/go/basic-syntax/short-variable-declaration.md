---
title: "ประกาศตัวแปรแบบสั้น"
description: "ตัวดำเนินการ := ประกาศและกำหนดค่าในขั้นตอนเดียว"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`:=` ประกาศตัวแปรพร้อมกำหนดค่า และอนุมาน type ให้ เป็นวิธีหลักที่ใช้ใน function

## ใช้ `:=`

```go title="short.go"
count := 10
name, ok := "go", true
count, extra := 20, 1 // ได้: extra เป็นตัวใหม่ ส่วน count แค่กำหนดค่าใหม่
fmt.Println(count, name, ok, extra) // 20 go true 1
```

ฝั่งซ้ายต้องมีตัวแปรใหม่อย่างน้อย 1 ตัว ตัวที่เหลือแค่กำหนดค่าใหม่ เลยเขียน `v, err := f()` ซ้ำได้หลายครั้งใน function เดียว

## `:=` vs `var`

| ใช้ | เมื่อ |
| --- | --- |
| `x := value` | อยู่ใน function และมีค่าอยู่แล้ว |
| `var x T` | ต้องการ zero value หรืออยู่ระดับ package |

:::caution[ข้อควรระวัง]
- `:=` ใน `if`/`for` สร้างตัวแปร **ใหม่** ที่ shadow ตัวนอก bug คลาสสิกคือ `err` ที่กำหนดในบล็อกไม่ถึง `err` ข้างนอก
- ใช้ `:=` ระดับ package ไม่ได้
:::
