---
title: "Type Constraints"
description: "interface ที่ระบุว่าใช้ type หรือ operation ไหนได้บ้าง"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

constraint คือ interface นอกจาก method แล้วยังระบุ type ที่อนุญาตด้วย `|` ได้ ทำให้ใช้ operator อย่าง `+` และ `<` ใน function ได้

## ตัวอย่าง

```go title="sum.go"
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n // ได้: ทุก type ใน Number รองรับ +
	}
	return total
}

type Score int
Sum([]int{1, 2})       // 3
Sum([]float64{1.5, 2}) // 3.5
Sum([]Score{3, 4})     // 7 เพราะมี ~int
```

## Constraint ในตัว

| Constraint | ทำได้ |
| --- | --- |
| `any` | ทุก type แต่แค่กำหนดค่าและส่งต่อ |
| `comparable` | `==`, `!=` และใช้เป็น key ของ map ได้ |
| `cmp.Ordered` | `< <= > >=` กับตัวเลขและ string |

:::caution[ข้อควรระวัง]
- interface ที่มีรายการ type (`int | float64`) ใช้ได้แค่เป็น constraint ใช้เป็น type ของตัวแปรธรรมดาไม่ได้
:::
