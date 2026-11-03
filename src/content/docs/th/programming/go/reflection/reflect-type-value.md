---
title: "reflect.Type และ reflect.Value"
description: "Type vs Kind และการแก้ค่าผ่าน pointer"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`Type` บอกว่า **เป็นอะไร** ส่วน `Value` ถือ **ข้อมูล** `Kind` คือหมวดพื้นฐาน (int, struct, slice …) ส่วน `Type` คือชื่อที่ประกาศไว้

## ตัวอย่าง

```go title="typevalue.go"
type Celsius float64

t := reflect.TypeOf(Celsius(1))
fmt.Println(t.Name(), t.Kind()) // Celsius float64

x := 1
v := reflect.ValueOf(&x).Elem() // ผ่าน pointer → แก้ได้
fmt.Println(v.CanSet()) // true
v.SetInt(42)
fmt.Println(x) // 42
```

## Method ที่มีประโยชน์

| บน `Type` | บน `Value` |
| --- | --- |
| `Name`, `Kind`, `String` | `Kind`, `Type`, `Interface()` |
| `NumField`, `Field(i)` | `Field(i)`, `FieldByName` |
| `Elem` (สิ่งที่ pointer/slice ชี้) | `Elem`, `Len`, `Index(i)` |
| `NumMethod`, `Method(i)` | `Set`, `SetInt`, `SetString` |

:::caution[ข้อควรระวัง]
- `reflect.ValueOf(x)` ถือสำเนา ถ้าจะแก้ต้องส่ง `&x` แล้วเรียก `.Elem()`
- field ที่ไม่ export อ่านได้แต่แก้ไม่ได้
:::
