---
title: "Variadic Functions"
description: "รับ argument กี่ตัวก็ได้ด้วย ...T"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

parameter ตัวสุดท้ายที่เป็น `...T` รับ argument ได้ตั้งแต่ 0 ตัวขึ้นไป ข้างใน function มันคือ `[]T`

## ประกาศและกระจาย slice

```go title="sum.go"
func sum(nums ...int) int { // nums เป็น []int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

sum()        // 0
sum(1, 2, 3) // 6
xs := []int{4, 5}
sum(xs...)   // 9, กระจาย slice
```

`fmt.Println(a ...any)` และ `append(s, items...)` คือ variadic function ที่ใช้ทุกวัน

:::caution[ข้อควรระวัง]
- variadic ได้เฉพาะ parameter ตัวสุดท้าย
- `f(xs...)` ส่ง slice ตัวเดิมไป ไม่ใช่สำเนา function แก้ element ของเราได้
:::
