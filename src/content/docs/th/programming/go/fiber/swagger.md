---
title: "Swagger"
description: "สร้าง OpenAPI doc จาก comment ด้วย swag"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`swaggo/swag` อ่าน comment ที่เป็น annotation บน handler แล้วสร้าง spec ของ OpenAPI (Swagger) ส่วน `gofiber/swagger` เปิดหน้า Swagger UI ให้

## ใส่ annotation ให้ handler

```go title="books.go"
// GetBook godoc
//
//	@Summary	Get a book by ID
//	@Tags		books
//	@Produce	json
//	@Param		id	path		int	true	"Book ID"
//	@Success	200	{object}	Book
//	@Failure	404	{object}	map[string]string
//	@Router		/books/{id} [get]
func GetBook(c *fiber.Ctx) error { ... }
```

ใส่ข้อมูลทั่วไป (`@title`, `@version`, `@BasePath`) ไว้เหนือ `main`

## Generate และเปิดหน้า

```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init            # เขียน docs/ (swagger.json, swagger.yaml, docs.go)
```

```go
import (
	_ "myapp/docs"
	"github.com/gofiber/swagger"
)

app.Get("/swagger/*", swagger.HandlerDefault) // UI อยู่ที่ /swagger/index.html
```

:::caution[ข้อควรระวัง]
- annotation เพี้ยนจากโค้ดได้แบบเงียบๆ ให้รัน `swag init` ใน CI (หรือผ่าน `go generate`) แล้วให้ fail ถ้ามี diff
:::
