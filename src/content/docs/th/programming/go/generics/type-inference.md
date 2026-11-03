---
title: "Type Inference"
description: "เมื่อ compiler หา type argument ให้เอง"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

compiler อนุมาน type argument จาก argument ของ function ให้ แทบไม่ต้องเขียนเอง

## ตัวอย่าง

```go title="infer.go"
Map([]int{1, 2}, strconv.Itoa) // อนุมาน T=int, U=string → [1 2]

func Zero[T any]() T {
	var z T
	return z
}
Zero[string]() // "" : ไม่มี argument เลยต้องระบุ T
Zero[int]()    // 0
```

## เมื่อไหร่ต้องระบุเอง

- type parameter โผล่แค่ในผลลัพธ์ (`Zero[T]`, constructor แบบ `make`)
- untyped constant จะอนุมานเป็น type ที่ไม่ต้องการ: `Larger[float64](1, 2)`

:::caution[ข้อควรระวัง]
- การระบุบางส่วนทำจากซ้ายไปขวา: `Map[int]` กำหนด `T` แล้ว `U` อนุมานเอง
:::
