---
title: "Generics Best Practices"
description: "เมื่อไหร่ generics ช่วย และเมื่อไหร่ interface ดีกว่า"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

ใช้ generics เมื่อโค้ดเดียวกันจะต้องถูก copy ไปใช้กับหลาย type ถ้า function แค่เรียก method ใช้ interface ธรรมดาง่ายกว่า

## Interface vs generic

```go title="choose.go"
// interface พอแล้ว: function แค่เรียก method
func countBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}

// generics เหมาะ: algorithm ต้องใช้ type ของ element เอง
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}
```

## หลักง่ายๆ

- **เขียนเวอร์ชัน concrete ก่อน** แล้วค่อยทำเป็น generic เมื่อมี type ที่สองต้องการ
- **เหมาะ:** container, helper ของ slice/map, wrapper ของ API ที่มี type
- **ไม่เหมาะ:** business logic, อะไรที่เป็น interface อยู่แล้ว, ชั้น "generic repository" ที่ครอบ ORM
- **constraint ให้น้อยที่สุด:** ขอแค่ `any` หรือ `comparable` ถ้าไม่ต้องใช้ operator

:::caution[ข้อควรระวัง]
- type parameter ทำให้ signature อ่านยากขึ้น อย่าเพิ่มเพื่อผู้เรียกแค่คนเดียว
:::
