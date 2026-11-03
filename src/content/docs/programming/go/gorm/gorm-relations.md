---
title: "GORM Relations"
description: "Has many, belongs to and Preload"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

GORM infers relations from struct fields and foreign keys. Associations are not loaded unless you ask with `Preload` or `Joins`.

## Example

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
db.Create(&Customer{Name: "ann", Orders: []Order{{Total: 10}, {Total: 25}}}) // saves orders too

var c Customer
db.Preload("Orders").First(&c, "name = ?", "ann") // len(c.Orders) == 2

var o Order
db.Preload("Customer").First(&o) // o.Customer.Name == "ann"
```

| Relation | Fields |
| --- | --- |
| Belongs to | `CustomerID uint` + `Customer Customer` |
| Has one | `Profile Profile` on the owner, `UserID` on Profile |
| Has many | `Orders []Order` on the owner, `CustomerID` on Order |
| Many to many | `Tags []Tag` with tag `gorm:"many2many:post_tags"` |

:::caution[Pitfalls]
- Loading associations in a loop (`for _, c := range customers { db.Model(&c).Association("Orders")... }`) is the N+1 problem; use `Preload` once.
:::
