---
title: "พื้นฐาน Map"
description: "ค้นหาแบบ key-value ด้วย map[K]V"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

map เก็บคู่ key-value และค้นหาด้วย key ได้เร็ว key ต้องเทียบด้วย `==` ได้ (string, ตัวเลข, struct ที่ประกอบจากพวกนั้น, array) ส่วน slice, map และ func เป็น key ไม่ได้

## สร้าง เขียน อ่าน

```go title="map.go"
ages := map[string]int{"ann": 30}
ages["bob"] = 25
fmt.Println(ages["ann"], ages["zed"], len(ages)) // 30 0 2
```

อ่าน key ที่ไม่มีจะได้ zero value ของ type ไม่ใช่ error

## Nil map

```go title="nil.go"
var m map[string]int // nil map: อ่านได้ เขียนแล้ว panic
fmt.Println(m["x"], m == nil) // 0 true
```

สร้าง map ด้วย literal หรือ `make(map[K]V)` ก่อนเขียน

:::caution[ข้อควรระวัง]
- `m["k"] = v` บน nil map จะ panic: "assignment to entry in nil map" เจอบ่อยกับ field ที่เป็น map ใน struct
- map ไม่ปลอดภัยถ้าเขียนพร้อมกันหลาย goroutine runtime อาจ crash ด้วย "concurrent map writes"
:::
