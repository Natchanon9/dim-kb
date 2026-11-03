---
title: "Map Operations"
description: "เพิ่ม แก้ ลบ นับ และล้าง"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

การทำงานกับ map เป็นส่วนหนึ่งของภาษา ไม่ใช่ method

## Operation ต่างๆ

```go title="ops.go"
m := map[string]int{"a": 1, "b": 2}
m["a"] = 10       // เพิ่มหรือแก้
delete(m, "b")    // ลบ
delete(m, "nope") // ไม่มี key ก็ไม่ error
fmt.Println(m, len(m)) // map[a:10] 1
clear(m)               // Go 1.21+: ลบทั้งหมด
```

| Operation | Syntax |
| --- | --- |
| เพิ่ม / แก้ | `m[k] = v` |
| อ่าน | `v := m[k]` หรือ `v, ok := m[k]` |
| ลบ | `delete(m, k)` |
| นับ | `len(m)` |
| ล้าง | `clear(m)` |

`fmt` พิมพ์ map โดยเรียง key ให้ output ของ test จึงคงที่

:::caution[ข้อควรระวัง]
- เอา address ของ element ไม่ได้: `&m[k]` compile ไม่ผ่าน และกับ `map[string]Point` จะเขียน `m[k].X = 1` ไม่ได้ ให้เก็บเป็น pointer หรือกำหนดค่าใหม่ทั้งก้อน
:::
