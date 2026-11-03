---
title: "Generic Functions"
description: "Filter, Reduce และ helper อื่นที่ใช้ซ้ำได้"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

generic function เด่นมากกับ algorithm เล็กๆ บน collection ที่ต่างกันแค่ type ของ element

## ตัวอย่าง

```go title="filter.go"
func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

evens := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }) // [2 4]
total := Reduce(evens, 0, func(acc, n int) int { return acc + n })        // 6
```

เช็ค package `slices` และ `maps` ของ standard library ก่อน helper หลายตัว (`slices.IndexFunc`, `slices.DeleteFunc`) มีให้แล้ว

:::caution[ข้อควรระวัง]
- การต่อ Filter/Map/Reduce เป็นสาย ไม่เป็นสไตล์ Go เท่า `for` loop ธรรมดา ใช้ helper เฉพาะตอนที่ทำให้โค้ดอ่านง่ายขึ้น
:::
