---
title: "Fiber Middleware"
description: "app.Use, c.Next และ middleware ของกลุ่ม route"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

middleware ของ Fiber คือ handler ที่เรียก `c.Next()` เพื่อทำต่อ หรือ return ก่อนเพื่อหยุด request

## ตัวอย่าง

```go title="middleware.go"
func requireAPIKey(c *fiber.Ctx) error {
	if c.Get("X-API-Key") != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.Next()
}

app := fiber.New()
app.Use(recover.New()) // มีในตัว: panic → 500
app.Use(func(c *fiber.Ctx) error {
	log.Println("->", c.Method(), c.Path())
	return c.Next()
})

api := app.Group("/api", requireAPIKey) // เฉพาะ /api/*
api.Get("/data", func(c *fiber.Ctx) error { return c.SendString("data") })
// ไม่มี key → 401 ; X-API-Key: secret → 200 data
```

## Middleware ในตัวที่ควรรู้

| Package | ทำอะไร |
| --- | --- |
| `middleware/logger` | access log |
| `middleware/recover` | กู้คืนจาก panic |
| `middleware/cors` | header ของ CORS |
| `middleware/limiter` | rate limiting |
| `middleware/requestid` | header ของ request ID |

:::caution[ข้อควรระวัง]
- ลำดับสำคัญ: `app.Use` รันตามลำดับที่ลงทะเบียน และต้องอยู่ก่อน route ที่ต้องการป้องกัน
:::
