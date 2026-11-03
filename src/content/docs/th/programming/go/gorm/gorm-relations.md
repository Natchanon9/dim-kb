---
title: "GORM Relations"
description: "has many, belongs to และ Preload"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

GORM อนุมานความสัมพันธ์จาก field ของ struct และ foreign key association จะไม่ถูกโหลดจนกว่าจะขอด้วย `Preload` หรือ `Joins`

## ตัวอย่าง

```go title="relations.go"
type Customer struct {
	gorm.Model
	Name   string
	Orders []Order // has many
}

type Order struct {
	gorm.Model
	CustomerID uint     // foreign key
	Customer   Customer // belongs to
	Total      int
}

db.AutoMigrate(&Customer{}, &Order{})
db.Create(&Customer{Name: "ann", Orders: []Order{{Total: 10}, {Total: 25}}}) // บันทึก order ไปด้วย

var c Customer
db.Preload("Orders").First(&c, "name = ?", "ann") // len(c.Orders) == 2

var o Order
db.Preload("Customer").First(&o) // o.Customer.Name == "ann"
```

| ความสัมพันธ์ | Field |
| --- | --- |
| Belongs to | `CustomerID uint` + `Customer Customer` |
| Has one | `Profile Profile` ที่ฝั่งเจ้าของ และ `UserID` ที่ Profile |
| Has many | `Orders []Order` ที่ฝั่งเจ้าของ และ `CustomerID` ที่ Order |
| Many to many | `Tags []Tag` พร้อม tag `gorm:"many2many:post_tags"` |

:::caution[ข้อควรระวัง]
- โหลด association ใน loop (`for _, c := range customers { db.Model(&c).Association("Orders")... }`) คือปัญหา N+1 ให้ใช้ `Preload` ครั้งเดียว
:::
