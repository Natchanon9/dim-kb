---
title: "If แบบมี Init"
description: "รัน statement สั้นๆ ก่อนเงื่อนไข และจำกัด scope ของตัวแปร"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`if` ขึ้นต้นด้วย statement ได้ โดยคั่นกับเงื่อนไขด้วย `;` ตัวแปรที่ประกาศตรงนั้นมีอยู่แค่ในบล็อก `if`/`else`

## รูปแบบ

```go title="init.go"
stock := map[string]int{"apple": 3}
if n, ok := stock["apple"]; ok {
	fmt.Println("in stock:", n) // in stock: 3
}
if _, err := strconv.Atoi("x"); err != nil {
	fmt.Println("error:", err)
}
// ตรงนี้ไม่มี n กับ err แล้ว
```

ใช้กับการเช็คแบบ comma-ok และ error ที่จัดการทันที ตัวแปรที่ใช้แป๊บเดียวจะไม่หลุดไปปนกับส่วนอื่นของ function

:::caution[ข้อควรระวัง]
- ถ้าต้องใช้ค่าหลัง `if` ให้ประกาศไว้ก่อน ไม่อย่างนั้นค่าหายไปแล้ว
- statement ใน init ควรสั้น ถ้าเรียก function ยาวๆ ตรงนั้นจะอ่านยาก
:::
