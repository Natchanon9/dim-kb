---
title: "พื้นฐาน Method"
description: "ผูก function กับ type ด้วย receiver"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

method คือ function ที่มี **receiver** เขียนไว้หน้าชื่อ ประกาศ method ได้กับทุก named type ที่ประกาศใน package ของเรา ไม่ใช่แค่ struct

## ประกาศ method

```go title="methods.go"
type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { return r.W * r.H }

type Celsius float64

func (c Celsius) Fahrenheit() float64 { return float64(c)*9/5 + 32 }

fmt.Println(Rect{3, 4}.Area(), Celsius(100).Fahrenheit()) // 12 212
```

ชื่อ receiver สั้นๆ ส่วนใหญ่ใช้ตัวอักษร 1–2 ตัวของชื่อ type (`r`, `c`) ไม่ใช้ `this` หรือ `self`

:::caution[ข้อควรระวัง]
- เพิ่ม method ให้ type ของ package อื่น (`time.Time`, `int`) ไม่ได้ ให้สร้าง named type ของตัวเองแทน
:::
