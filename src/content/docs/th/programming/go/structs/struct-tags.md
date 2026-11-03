---
title: "Struct Tags"
description: "metadata ของ field สำหรับ json, db และ library ตรวจข้อมูล"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

tag คือ string ที่แปะไว้กับ field package ต่างๆ อ่านผ่าน reflection เพื่อตัดสินใจว่าจะ encode, เก็บ หรือ validate field นั้นอย่างไร

## ตัวอย่าง

```go title="tags.go"
type Item struct {
	Name   string `json:"name"`
	Price  int    `json:"price,omitempty"`
	Secret string `json:"-"`
}

b, _ := json.Marshal(Item{Name: "pen", Secret: "x"})
f, _ := reflect.TypeOf(Item{}).FieldByName("Name")
fmt.Println(string(b), f.Tag.Get("json")) // {"name":"pen"} name
```

## Tag ที่เจอบ่อย

| Tag | ใช้กับ |
| --- | --- |
| `json:"name,omitempty"` | `encoding/json` |
| `db:"user_id"` | sqlx และตัวอื่นที่คล้ายกัน |
| `gorm:"primaryKey"` | GORM |
| `validate:"required,email"` | go-playground/validator |

ใส่หลาย tag ใน string เดียวได้โดยคั่นด้วยช่องว่าง: `` `json:"id" db:"id"` ``

:::caution[ข้อควรระวัง]
- รูปแบบเข้มงวด: `key:"value"` ห้ามมีช่องว่างรอบ colon พิมพ์ผิดจะเงียบไม่ทำอะไร แต่ `go vet` จับ tag ที่ผิดรูปแบบได้
:::
