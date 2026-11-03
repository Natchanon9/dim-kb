---
title: "ตัวแปรและการประกาศ"
description: "ประกาศตัวแปรด้วย var, type inference และการประกาศเป็นกลุ่ม"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`var` ใช้ประกาศตัวแปร โดยระบุชื่อกับ type หรือค่า หรือทั้งคู่ Go เป็นภาษา statically typed type ของตัวแปรเปลี่ยนไม่ได้

## `var` 3 แบบ

```go title="vars.go"
var name string = "gopher"
var age = 3 // อนุมาน type เป็น int
var (
	host = "localhost"
	port = 8080
)
fmt.Println(name, age, host, port) // gopher 3 localhost 8080
```

- **type + ค่า:** ใช้เมื่อ type ที่อนุมานได้ไม่ใช่ที่ต้องการ (`var ratio float64 = 1`)
- **ค่าอย่างเดียว:** type มาจากค่า
- **เป็นกลุ่ม `var (...)`:** นิยมใช้กับตัวแปรระดับ package

## Scope

ตัวแปรมองเห็นได้ภายในบล็อก `{}` ที่ประกาศ ตัวแปรระดับ package มองเห็นทั้ง package และใช้ `:=` ไม่ได้

:::caution[ข้อควรระวัง]
- ตัวแปร local ที่ไม่ได้ใช้ compile ไม่ผ่าน
- ประกาศชื่อเดิมในบล็อกด้านในจะ **shadow** ตัวนอก แก้ค่าข้างในแล้วตัวนอกไม่เปลี่ยน
:::
