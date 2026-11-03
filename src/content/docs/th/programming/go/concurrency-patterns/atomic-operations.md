---
title: "Atomic Operations"
description: "ตัวนับและ flag แบบไม่ใช้ lock ด้วย sync/atomic"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`sync/atomic` อัปเดตค่าตัวเดียวโดยไม่ต้อง lock เหมาะกับตัวนับ flag และการสลับ pointer ไปยัง config ที่ไม่ถูกแก้

## ตัวอย่าง

```go title="atomic.go"
var hits atomic.Int64
// goroutine 1000 ตัวเรียก hits.Add(1)
fmt.Println(hits.Load()) // 1000

var state atomic.Int32
state.CompareAndSwap(0, 1) // true: 0 → 1
state.CompareAndSwap(0, 1) // false: เป็น 1 แล้ว
```

atomic แบบมี type (`atomic.Int64`, `atomic.Bool`, `atomic.Pointer[T]` ตั้งแต่ Go 1.19) ปลอดภัยกว่า function แบบเก่าอย่าง `atomic.AddInt64(&x, 1)`

## Atomic หรือ mutex?

ค่าเดียวที่เป็นอิสระ → atomic ส่วนหลายค่าที่ต้องเปลี่ยนพร้อมกัน → mutex

:::caution[ข้อควรระวัง]
- ผสมการเข้าถึงแบบ atomic กับแบบธรรมดาบนตัวแปรเดียวกันก็ยังเป็น data race
:::
