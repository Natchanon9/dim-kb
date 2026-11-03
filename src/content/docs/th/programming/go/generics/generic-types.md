---
title: "Generic Types"
description: "struct และ type อื่นที่มี type parameter"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

type รับ type parameter ได้เหมือนกัน method จะใช้ parameter ของ type แต่เพิ่มตัวใหม่เองไม่ได้

## ตัวอย่าง: Stack

```go title="stack.go"
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

var s Stack[string] // ระบุ concrete type
s.Push("a")
s.Push("b")
s.Pop() // b true
```

เหมาะกับ: container (stack, queue, set, LRU cache) และ wrapper ที่มี type (`Result[T]`, `Page[T]` สำหรับ API response แบบแบ่งหน้า)

:::caution[ข้อควรระวัง]
- `var zero T` คือวิธีคืนค่า "ว่าง" ของ type ที่ไม่รู้ล่วงหน้า
:::
