---
title: "GORM Basics"
description: "Models, AutoMigrate and CRUD with GORM"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

GORM is the most used Go ORM: structs are tables, and methods build SQL for you.

## Setup

```sh
go get gorm.io/gorm gorm.io/driver/postgres
```

```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

## Model and CRUD

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
db.Delete(&p)                       // soft delete: sets deleted_at

err := db.First(&Product{}, p.ID).Error
errors.Is(err, gorm.ErrRecordNotFound) // true
```

Use it from Fiber handlers through a repository, passing `db.WithContext(c.UserContext())` so queries are cancelled with the request.

:::caution[Pitfalls]
- Always check `.Error` on the result; GORM doesn't return errors as a second value.
- `Updates(struct)` skips zero fields; use a `map[string]any` or `Select` to set a field to 0 or "".
- `AutoMigrate` never drops columns; use real migrations (goose, golang-migrate) in production.
:::
