---
title: "พื้นฐาน Pointer"
description: "& เอา address ส่วน * อ่านหรือเขียนผ่าน address"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

pointer เก็บ address ในหน่วยความจำของค่า `&x` เอา address ของ x ส่วน `*p` อ่านหรือเขียนค่าที่ p ชี้อยู่

## ตัวอย่าง

```go title="pointer.go"
x := 10
p := &x // p เป็น *int ที่เก็บ address ของ x
*p = 20 // เขียนผ่าน pointer
fmt.Println(x, *p, p != nil) // 20 20 true
```

| Syntax | ความหมาย |
| --- | --- |
| `*T` | type "pointer ไปยัง T" |
| `&x` | address ของ x |
| `*p` | ค่าที่ p ชี้ (dereference) |
| `nil` | zero value ของ pointer ทุกตัว |

Go ไม่มี pointer arithmetic (นอกจากใน `unsafe`) และ garbage collector จะคืน memory เมื่อไม่มีอะไรชี้ไปที่ค่านั้นแล้ว

:::caution[ข้อควรระวัง]
- dereference pointer ที่เป็น nil จะ panic: "invalid memory address or nil pointer dereference"
:::
