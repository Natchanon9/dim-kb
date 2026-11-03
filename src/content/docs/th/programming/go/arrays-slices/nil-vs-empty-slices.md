---
title: "Nil vs Empty Slice"
description: "ยาว 0 เหมือนกัน แต่ต่างกันตอนเทียบกับ nil และใน JSON"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

nil slice กับ empty slice ยาว 0 ทั้งคู่ และใช้กับ `len`, `range`, `append` ได้เหมือนกัน แต่ต่างกันอยู่ 2 จุด

## ความต่าง

```go title="nil.go"
var a []int  // nil
b := []int{} // ว่าง แต่ไม่ใช่ nil
fmt.Println(a == nil, b == nil, len(a), len(b)) // true false 0 0

ja, _ := json.Marshal(a) // null
jb, _ := json.Marshal(b) // []
```

## ควรใช้แบบไหน

- ในโค้ดทั่วไปใช้ `var s []T` และเช็คด้วย `len(s) == 0` อย่าเช็ค `s == nil`
- ใน API response ที่ client คาดว่าจะได้ `[]` ให้ initialize ด้วย `[]T{}` (หรือ `make([]T, 0)`)

:::caution[ข้อควรระวัง]
- field ใน JSON ที่บางทีเป็น `null` บางทีเป็น `[]` ทำให้ client ที่เข้มงวดพัง เลือกแบบเดียวให้ชัด
:::
