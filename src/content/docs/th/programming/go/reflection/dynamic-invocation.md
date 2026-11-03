---
title: "Dynamic Invocation"
description: "เรียก method และ function ตามชื่อตอน runtime"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`Value.MethodByName` หา method ตอน runtime และ `Value.Call` เรียกมันด้วย argument ที่เป็น `[]reflect.Value`

## ตัวอย่าง

```go title="call.go"
type Calc struct{}

func (Calc) Add(a, b int) int { return a + b }

m := reflect.ValueOf(Calc{}).MethodByName("Add")
out := m.Call([]reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)})
fmt.Println(out[0].Int()) // 5
```

ใช้ใน RPC framework, ระบบ plugin และ router ของคำสั่งที่จับคู่ชื่อกับ method

## ส่วนใหญ่มีทางที่ดีกว่า

`map[string]func(...)` ให้การเลือก function แบบเดียวกัน แต่เช็คได้ตอน compile และไม่ต้องใช้ reflection

:::caution[ข้อควรระวัง]
- `MethodByName` คืน `Value` ที่ใช้ไม่ได้ถ้าไม่มี method นั้นหรือไม่ได้ export ให้เช็ค `m.IsValid()` ก่อน `Call`
- จำนวนหรือ type ของ argument ผิดจะ panic
:::
