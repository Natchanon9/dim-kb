---
title: "Anonymous Functions"
description: "function literal ที่ไม่มีชื่อ"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

function literal ไม่มีชื่อ เขียนแทรกไว้ตรงที่ต้องการค่า function หรือเรียกทันทีก็ได้

## ที่ใช้บ่อย

```go title="anon.go"
result := func(a, b int) int { return a * b }(3, 4) // ประกาศแล้วเรียกเลย: 12

names := []string{"bob", "al", "christine"}
sort.Slice(names, func(i, j int) bool { return len(names[i]) < len(names[j]) })
// [al bob christine]
```

จะได้เขียนบ่อยใน comparator ของ sort, `defer func() { ... }()`, `go func() { ... }()` และ HTTP handler

:::caution[ข้อควรระวัง]
- anonymous function ที่ใช้ตัวแปรข้างนอกคือ closure จะเห็นค่า **ปัจจุบัน** ตอนที่รัน ไม่ใช่ค่าตอนที่สร้าง
- literal ที่ยาวทำให้อ่านยาก ยาวเกินไม่กี่บรรทัดก็ตั้งชื่อเป็น function แยกเถอะ
:::
