---
title: "range over func"
description: "เขียน iterator ของเราเองที่ใช้กับ for-range ได้ (Go 1.23+)"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

ตั้งแต่ Go 1.23 `for range` รับ function ได้ function นั้นเรียก `yield` ทีละค่า และ `yield` คืน false เมื่อ loop หยุดก่อน

## ตัวอย่าง

```go title="countdown.go"
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i > 0; i-- {
			if !yield(i) { // body ของ loop เจอ break/return
				return
			}
		}
	}
}

for v := range Countdown(5) {
	if v == 2 {
		break
	}
	fmt.Println(v) // 5 4 3
}
```

## รูปแบบที่รับได้

| Type ของ function | รูปแบบ loop |
| --- | --- |
| `func(yield func() bool)` | `for range f` |
| `func(yield func(V) bool)` | `for v := range f` |
| `func(yield func(K, V) bool)` | `for k, v := range f` |

:::caution[ข้อควรระวัง]
- เช็คผลของ `yield` เสมอ เรียกซ้ำหลังมันคืน false จะ panic
- ต้องมี `go 1.23` ใน `go.mod` (หรือไฟล์ที่มี `//go:build go1.23`)
:::
