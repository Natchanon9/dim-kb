---
title: "Fiber CRUD"
description: "Build a REST API with Fiber: routes, params, body parsing"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Fiber is an Express-style web framework built on fasthttp. Handlers take a `*fiber.Ctx` and return an `error`.

## Setup

```sh
go get github.com/gofiber/fiber/v2
```

## CRUD for books

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

| Need | Fiber |
| --- | --- |
| Path param | `c.Params("id")`, `c.ParamsInt("id")` |
| Query | `c.Query("page")` |
| Body | `c.BodyParser(&v)` |
| Respond | `c.JSON(v)`, `c.Status(201).JSON(v)`, `c.SendStatus(204)` |

Test without a port: `app.Test(httptest.NewRequest(...))`.

:::caution[Pitfalls]
- Values from `c.Params`/`c.Body` are only valid during the handler (fasthttp reuses buffers); copy them before keeping them.
- Fiber isn't `net/http`, so standard middleware doesn't plug in directly.
:::
