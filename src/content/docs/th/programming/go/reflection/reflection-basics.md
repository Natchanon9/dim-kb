---
title: "พื้นฐาน Reflection"
description: "ตรวจ type และค่าตอน runtime ด้วย reflect"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

reflection ให้โปรแกรมดู type และเนื้อหาของค่าได้ตอน runtime เป็นวิธีที่ `encoding/json`, `fmt` และ ORM ใช้จัดการ type ที่ไม่เคยรู้จักมาก่อน

## ตัวอย่าง

```go title="reflect.go"
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println(v.Type(), v.Kind(), v.Float()) // float64 float64 3.4
```

## ทางเข้าสองทาง

| Function | ได้อะไร |
| --- | --- |
| `reflect.TypeOf(x)` | `reflect.Type`: ชื่อ, kind, field, method |
| `reflect.ValueOf(x)` | `reflect.Value`: ข้อมูล พร้อม getter และ setter |

## ควรใช้ไหม?

ในโค้ดแอปแทบไม่ต้องใช้ ลอง interface หรือ generics ก่อน reflection เหมาะกับ library: encoder, validator, dependency injection, helper ของ test

:::caution[ข้อควรระวัง]
- เรียก getter ผิดตัว (`v.Int()` กับ string) จะ panic ตอน runtime ให้เช็ค `v.Kind()` ก่อน
:::
