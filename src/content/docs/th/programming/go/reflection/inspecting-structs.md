---
title: "ตรวจดู Struct"
description: "วน loop ทุก field ของ struct ด้วย reflection"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`Type.NumField` กับ `Type.Field(i)` ไล่ดู field ของ struct ได้: ชื่อ, type, tag และ export หรือไม่

## ตัวอย่าง

```go title="fields.go"
type Account struct {
	Name  string
	Age   int
	email string
}

t := reflect.TypeOf(Account{})
for i := range t.NumField() {
	f := t.Field(i)
	fmt.Println(f.Name, f.Type, f.IsExported())
}
// Name string true
// Age int true
// email string false
```

ใช้ `reflect.ValueOf(s).Field(i)` คู่กันเพื่ออ่านค่าของ field นั้น field ที่ embed มาจะมี `f.Anonymous == true`

:::caution[ข้อควรระวัง]
- ถ้าถือ pointer อยู่ให้เรียก `.Elem()` ก่อน `NumField` บน type ที่เป็น pointer จะ panic
:::
