---
title: "จัดการค่า NULL"
description: "scan column ที่เป็น NULL ด้วย type sql.Null หรือ pointer"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

scan `NULL` ของ SQL ลง `string` หรือ `int` ธรรมดาจะพัง ต้องใช้ type ที่บอกได้ว่า "ไม่มีค่า"

## ตัวอย่าง

```go title="null.go"
var phone sql.NullString
phone.Scan(nil) // แบบเดียวกับที่ rows.Scan ทำเมื่อเจอ NULL
fmt.Println(phone.Valid, phone.String == "") // false true

var age sql.Null[int64] // เวอร์ชัน generic, Go 1.22+
age.Scan(int64(30))
fmt.Println(age.Valid, age.V) // true 30

var nick *string // pointer: nil หมายถึง NULL
```

| ทางเลือก | ข้อดี |
| --- | --- |
| `sql.NullString`, `sql.NullInt64`, … | ชัดเจน ไม่ต้องจอง memory |
| `sql.Null[T]` (1.22+) | type generic ตัวเดียวใช้ได้หมด |
| `*T` | เข้ากับ JSON (`null`) และง่าย |
| `COALESCE(col, '')` ใน SQL | ใช้ type ธรรมดาได้ถ้า NULL = ค่าว่าง |

:::caution[ข้อควรระวัง]
- error "converting NULL to string is unsupported" แปลว่า scan column ที่เป็น NULL ได้ลง type ธรรมดา
:::
