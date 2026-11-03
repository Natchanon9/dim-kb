---
title: "Fiber CRUD"
description: "สร้าง REST API ด้วย Fiber: route, param, การอ่าน body"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Fiber คือ web framework สไตล์ Express ที่สร้างบน fasthttp handler รับ `*fiber.Ctx` และคืน `error`

## Setup

```sh
go get github.com/gofiber/fiber/v2
```

## CRUD ของหนังสือ

```go title="books.go"
app := fiber.New()

app.Post("/books", func(c *fiber.Ctx) error {
	var b Book
	if err := c.BodyParser(&b); err != nil || b.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title required"})
	}
	b.ID = store.Add(b)
	return c.Status(fiber.StatusCreated).JSON(b)
})

app.Get("/books/:id", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	b, ok := store.Get(id)
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(b)
})

app.Delete("/books/:id", func(c *fiber.Ctx) error { /* ... */ return c.SendStatus(fiber.StatusNoContent) })

app.Listen(":8080")
// POST {"title":"Go"} → 201 {"id":1,"title":"Go"}
// GET /books/1      → 200 {"id":1,"title":"Go"}
// GET /books/9      → 404 {"error":"not found"}
```

| ต้องการ | Fiber |
| --- | --- |
| path param | `c.Params("id")`, `c.ParamsInt("id")` |
| query | `c.Query("page")` |
| body | `c.BodyParser(&v)` |
| ตอบกลับ | `c.JSON(v)`, `c.Status(201).JSON(v)`, `c.SendStatus(204)` |

test โดยไม่ต้องเปิด port: `app.Test(httptest.NewRequest(...))`

:::caution[ข้อควรระวัง]
- ค่าจาก `c.Params`/`c.Body` ใช้ได้แค่ระหว่างที่ handler ทำงาน (fasthttp ใช้ buffer ซ้ำ) ต้อง copy ก่อนถ้าจะเก็บไว้
- Fiber ไม่ใช่ `net/http` middleware มาตรฐานจึงต่อเข้าตรงๆ ไม่ได้
:::
