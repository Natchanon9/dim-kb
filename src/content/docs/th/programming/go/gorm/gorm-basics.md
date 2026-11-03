---
title: "พื้นฐาน GORM"
description: "model, AutoMigrate และ CRUD ด้วย GORM"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

GORM คือ ORM ที่คนใช้มากที่สุดใน Go: struct คือ table และ method สร้าง SQL ให้

## Setup

```sh
go get gorm.io/gorm gorm.io/driver/postgres
```

```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

## Model และ CRUD

```go title="product.go"
type Product struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Code       string `gorm:"uniqueIndex"`
	Price      int
}

db.AutoMigrate(&Product{})

db.Create(&Product{Code: "P1", Price: 100})

var p Product
db.First(&p, "code = ?", "P1")      // 1 P1 100
db.Model(&p).Update("Price", 200)   // UPDATE ... SET price=200
db.Delete(&p)                       // soft delete: ตั้งค่า deleted_at

err := db.First(&Product{}, p.ID).Error
errors.Is(err, gorm.ErrRecordNotFound) // true
```

ใช้จาก Fiber handler ผ่าน repository และส่ง `db.WithContext(c.UserContext())` เพื่อให้ query ถูกยกเลิกพร้อม request

:::caution[ข้อควรระวัง]
- เช็ค `.Error` ของผลลัพธ์เสมอ GORM ไม่ได้คืน error เป็นค่าที่สอง
- `Updates(struct)` ข้าม field ที่เป็นศูนย์ ถ้าจะตั้งเป็น 0 หรือ "" ให้ใช้ `map[string]any` หรือ `Select`
- `AutoMigrate` ไม่ลบ column ใน production ให้ใช้ migration จริง (goose, golang-migrate)
:::
