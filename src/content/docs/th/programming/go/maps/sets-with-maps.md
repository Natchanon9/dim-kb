---
title: "ทำ Set ด้วย Map"
description: "ใช้ map[T]struct{} เป็น set"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Go ไม่มี type set ใช้ map ที่ value เป็น empty struct แทนได้ และ `struct{}` ใช้พื้นที่ 0 byte

## ตัวอย่าง

```go title="set.go"
seen := map[string]struct{}{}
for _, w := range []string{"go", "rust", "go"} {
	seen[w] = struct{}{}
}
_, ok := seen["go"]
fmt.Println(len(seen), ok) // 2 true
```

## `struct{}` vs `bool`

| Type ของ value | ข้อดี |
| --- | --- |
| `map[T]struct{}` | ไม่เปลือง memory ต่อค่า และเห็นชัดว่าเป็น set |
| `map[T]bool` | เช็คง่ายกว่า: `if seen[w] { ... }` |

ใช้ได้ทั้งคู่ ทีมส่วนใหญ่เลือกแบบเดียวให้เหมือนกันทั้งโปรเจกต์

:::caution[ข้อควรระวัง]
- union, intersection และ difference ต้องเขียน loop เอง (หรือทำ generic helper เล็กๆ)
:::
