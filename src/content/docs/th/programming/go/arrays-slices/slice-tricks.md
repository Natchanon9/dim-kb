---
title: "Slice Tricks"
description: "ลบ แทรก กลับลำดับ และกรองแบบ in place"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

การแก้ slice ที่ทำบ่อยเขียนบรรทัดเดียวได้ด้วย package `slices` (Go 1.21+) ส่วนการกรองแบบ in place ใช้ backing array เดิมเพื่อไม่ต้องจอง memory ใหม่

## ลบ แทรก กลับลำดับ

```go title="tricks.go"
s := []int{1, 2, 3, 4, 5}
s = slices.Delete(s, 1, 2) // ลบ index 1 → [1 3 4 5]
s = slices.Insert(s, 1, 2) // ใส่ 2 กลับ  → [1 2 3 4 5]
slices.Reverse(s)          //            → [5 4 3 2 1]
```

## กรองแบบ in place

```go title="filter.go"
evens := s[:0] // backing array เดิม ยาว 0
for _, v := range s {
	if v%2 == 0 {
		evens = append(evens, v)
	}
}
// evens → [4 2]
```

หรือใช้ `slices.DeleteFunc(s, func(v int) bool { return v%2 != 0 })`

:::caution[ข้อควรระวัง]
- กรองแบบ in place จะเขียนทับ slice เดิม หลังจากนั้นอย่าใช้ `s` อีก ใช้แค่ผลลัพธ์
- `Delete` กับ `Insert` คืน header ใหม่ ต้องรับค่ากลับเสมอ
:::
