---
title: "Comma-ok Idiom"
description: "แยก key ที่ไม่มี ออกจาก zero value"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`v, ok := m[k]` คืน boolean ตัวที่สองบอกว่ามี key นี้หรือไม่ ใช้ทุกครั้งที่ zero value เป็นค่าที่เป็นไปได้จริง

## ตัวอย่าง

```go title="ok.go"
stock := map[string]int{"apple": 0}
v, ok := stock["apple"] // 0 true  → มี และค่าเป็น 0
v, ok = stock["pear"]   // 0 false → ไม่มี
```

ส่วนใหญ่ใช้คู่กับ if แบบมี init:

```go
if price, ok := prices[item]; ok {
	total += price
}
```

รูปแบบ `v, ok` แบบเดียวกันนี้เจอใน type assertion (`x.(T)`) และการรับจาก channel (`v, ok := <-ch`)

:::caution[ข้อควรระวัง]
- ถ้าไม่ใช้ `ok` จะแยกไม่ออกว่า "ของเหลือ 0" หรือ "ไม่มีสินค้านี้"
:::
