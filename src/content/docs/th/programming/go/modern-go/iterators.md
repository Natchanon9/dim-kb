---
title: "Iterators"
description: "iter.Seq, iter.Seq2 และ iter.Pull"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

package `iter` (Go 1.23+) ตั้งชื่อ type ของ iterator: `iter.Seq[V]` ให้ค่าทีละตัว ส่วน `iter.Seq2[K, V]` ให้ทีละคู่ package มาตรฐานคืน type เหล่านี้แทนการสร้าง slice

## Seq2 และ Pull

```go title="iter.go"
func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

for i, s := range Enumerate([]string{"a", "b"}) { ... } // 0 a, 1 b

next, stop := iter.Pull(Countdown(3)) // ดึงค่าทีละตัว
defer stop()
v1, _ := next() // 3
v2, _ := next() // 2
```

## ทำไมต้อง iterator

- lazy: สร้างค่าเฉพาะตอนที่ถูกใช้ ลำดับไม่รู้จบหรือใหญ่มากก็ไม่เป็นไร
- ไม่ต้องมี slice ตรงกลาง
- รูปแบบมาตรฐานเดียวกันทุก library (`maps.Keys`, `slices.Values`, `strings.Lines` ใน 1.24)

:::caution[ข้อควรระวัง]
- หลัง `iter.Pull` ต้อง `defer stop()` เสมอ ไม่อย่างนั้น resource ของ iterator จะรั่ว
:::
