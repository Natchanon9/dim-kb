---
title: "Type Sets"
description: "~T หมายถึงอะไร และ constraint กำหนดชุดของ type อย่างไร"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

ทุก constraint อธิบาย **type set** คือชุดของ type ที่เข้ากับมัน `int` หมายถึง `int` เท่านั้น ส่วน `~int` หมายถึงทุก type ที่ underlying type เป็น `int`

## ตัวอย่าง

```go title="avg.go"
type Celsius float64

type Float interface {
	~float32 | ~float64 // ~ รวม named type อย่าง Celsius ด้วย
}

func Avg[T Float](xs ...T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total / T(len(xs))
}

Avg(Celsius(20), Celsius(30)) // 25
```

ถ้าไม่มี `~` จะใช้ `Celsius` ไม่ได้ ทั้งที่ข้างในเป็น `float64`

## การรวม

- `A | B`: union ของสองชุด
- หลายบรรทัดใน interface เดียวคือ intersection: `interface{ ~int | ~string; String() string }` = type ที่พื้นฐานเป็น int หรือ string และมี `String()` ด้วย

:::caution[ข้อควรระวัง]
- constraint ที่จะใช้ซ้ำควรใส่ `~` ไม่อย่างนั้น type ที่ผู้ใช้สร้างเองอย่าง `type UserID int` จะใช้ไม่ได้
:::
