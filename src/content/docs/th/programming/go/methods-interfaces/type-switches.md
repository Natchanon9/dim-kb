---
title: "Type Switches"
description: "จัดการหลาย concrete type ที่เก็บใน interface"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

type switch คือ type assertion หลายตัวต่อกันใน statement เดียว แต่ละ case ได้ตัวแปรเป็น type ของ case นั้น

## ตัวอย่าง

```go title="area.go"
func area(s any) float64 {
	switch v := s.(type) {
	case Rect:
		return v.Area()
	case *Square:
		return v.Area()
	default:
		return 0
	}
}
// area(Rect{2, 3}), area(&Square{3}), area("x") → 6 9 0
```

`case Rect` กับ `case *Square` ต่างกัน: ค่ากับ pointer เป็นคนละ type

:::caution[ข้อควรระวัง]
- type switch บน type ของเราเองมักแปลว่าขาด method ใน interface ถ้าทุก case เรียก method เดียวกัน ให้เรียกผ่าน interface แทน
:::
