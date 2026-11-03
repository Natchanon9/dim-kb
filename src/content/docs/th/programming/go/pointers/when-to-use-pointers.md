---
title: "ควรใช้ Pointer เมื่อไหร่"
description: "แก้ค่า, struct ใหญ่, state ที่แชร์ และค่าที่อาจไม่มี"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

ใช้ pointer เพราะมีเหตุผล ไม่ใช่เพราะความเคยชิน ค่าธรรมดาง่ายกว่าและมักเร็วกว่า

## ใช้ pointer เมื่อ

- **function ปลายทางต้องแก้ค่า**
- **struct ใหญ่** (field เยอะ, array ใหญ่) copy แล้วเปลือง
- **ค่าต้องถูกแชร์** หรือมี lock อยู่ข้างใน (`sync.Mutex`)
- **ต้องแยก "ไม่มีค่า"** ออกจาก zero value:

```go title="patch.go"
// "ไม่ได้ส่งมา" (nil) ต่างจาก "ตั้งเป็นศูนย์"
type Patch struct {
	Name *string
	Age  *int
}

age := 0
patch := Patch{Age: &age}
fmt.Println(patch.Name == nil, *patch.Age) // true 0
```

## ใช้ค่าธรรมดาเมื่อ

- ข้อมูลเล็กที่ไม่ถูกแก้ (`time.Time`, จุดพิกัด, จำนวนเงิน)
- map, slice, channel และ func: ชี้ไปที่ข้อมูลที่แชร์อยู่แล้ว

:::caution[ข้อควรระวัง]
- `[]*T` vs `[]T`: pointer ใน slice เพิ่มการจอง memory และงานของ GC ใช้ `[]T` ยกเว้น element ต้องถูกแชร์หรือใหญ่มาก
:::
