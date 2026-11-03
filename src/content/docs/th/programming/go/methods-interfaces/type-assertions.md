---
title: "Type Assertions"
description: "ดึงค่า concrete ออกจาก interface ด้วย x.(T)"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`x.(T)` ดึงค่าที่เป็น type `T` ออกจาก interface `x` แบบคืนสองค่าจะบอกว่าสำเร็จไหมแทนการ panic

## 2 รูปแบบ

```go title="assert.go"
var v any = "go"
s := v.(string)  // panic ถ้า v ไม่ใช่ string
n, ok := v.(int) // แบบปลอดภัย: ok บอกว่าสำเร็จไหม
fmt.Println(s, n, ok) // go 0 false
```

## Assert เป็น interface

`T` เป็น interface อื่นได้: `if s, ok := v.(fmt.Stringer); ok { ... }` เช็คว่าค่านั้นมี method `String` ไหม library มาตรฐานใช้วิธีนี้ตรวจความสามารถเสริม (เช่น `io.WriterTo`)

:::caution[ข้อควรระวัง]
- ใช้แบบคืนค่าเดียวเฉพาะตอนที่ type ผิดหมายถึง bug ของโปรแกรม นอกนั้นใช้ `, ok`
- กับ error ให้ใช้ `errors.As` แทนการ assert จะได้หา error ที่ถูก wrap เจอ
:::
