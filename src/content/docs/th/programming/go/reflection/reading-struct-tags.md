---
title: "อ่าน Struct Tags"
description: "สร้างการตรวจข้อมูลและการจับคู่จาก tag"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`StructField.Tag.Get("key")` อ่านค่าของ tag นี่คือกลไกทั้งหมดที่อยู่เบื้องหลัง tag `json`, `db` และ `validate`

## Validator ตัวเล็ก

```go title="validate.go"
type Signup struct {
	Email string `validate:"required"`
	Nick  string
}

func missingRequired(v any) []string {
	val := reflect.ValueOf(v)
	t := val.Type()
	var missing []string
	for i := range t.NumField() {
		if t.Field(i).Tag.Get("validate") == "required" && val.Field(i).IsZero() {
			missing = append(missing, t.Field(i).Name)
		}
	}
	return missing
}

missingRequired(Signup{Nick: "x"}) // [Email]
```

`Tag.Lookup(key)` บอกได้ด้วยว่ามี key นั้นอยู่ไหม

:::caution[ข้อควรระวัง]
- ใน production ใช้ library ที่มีคนดูแลอย่าง `go-playground/validator` ดีกว่าเขียนเอง
:::
