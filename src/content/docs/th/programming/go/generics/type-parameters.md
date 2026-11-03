---
title: "Type Parameters"
description: "ประกาศ [T Constraint] และใช้กับ concrete type"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

รายการ type parameter `[T Constraint]` อยู่หลังชื่อ function แต่ละตัวมี constraint บอกว่า type นั้นต้องรองรับอะไร

## ตัวอย่าง

```go title="larger.go"
func Larger[T cmp.Ordered](a, b T) T { // T ต้องรองรับ < และ >
	if a > b {
		return a
	}
	return b
}

Larger(3, 7)             // 7 อนุมาน T เป็น int
Larger("a", "b")         // b
Larger[float64](1, 2.5)  // 2.5 ระบุ T เอง
```

หลาย parameter: `[K comparable, V any]` ส่วน parameter ที่ใช้ constraint เดียวกัน: `[T, U any]`

:::caution[ข้อควรระวัง]
- argument ทั้งสองต้องเป็น `T` **ตัวเดียวกัน**: `Larger(1, 2.5)` อนุมานเป็น `float64` จาก untyped constant ได้ แต่ `Larger(x, y)` ที่ตัวแปรเป็น `int` กับ `float64` จะไม่ผ่าน
:::
