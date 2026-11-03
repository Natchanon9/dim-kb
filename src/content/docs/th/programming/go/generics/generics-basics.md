---
title: "พื้นฐาน Generics"
description: "เขียน function หรือ type เดียวที่ใช้ได้กับหลาย type (Go 1.18+)"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

generics (Go 1.18+) ให้ function หรือ type รับ **type parameter** ในวงเล็บเหลี่ยม เขียน algorithm ครั้งเดียว แล้ว compiler เช็คให้ทุก type ที่นำไปใช้

## ตัวอย่าง

```go title="map.go"
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) })
// ["2" "4" "6"]
```

## ก่อนมี generics

ต้อง copy function แยกต่อ type หรือใช้ `any` กับ type assertion ซึ่งเสียการเช็คตอน compile generics ให้ทั้งการใช้ซ้ำและความปลอดภัยของ type

:::caution[ข้อควรระวัง]
- method มี type parameter ของตัวเองไม่ได้ มีได้แค่ที่ type
:::
