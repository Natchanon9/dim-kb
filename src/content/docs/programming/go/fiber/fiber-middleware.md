---
title: "Fiber Middleware"
description: "app.Use, c.Next and route-group middleware"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Fiber middleware is a handler that calls `c.Next()` to continue, or returns early to stop the request.

## Example

```go title="middleware.go"
func requireAPIKey(c *fiber.Ctx) error {
	if c.Get("X-API-Key") != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.Next()
}

app := fiber.New()
app.Use(recover.New()) // built-in: panic → 500
app.Use(func(c *fiber.Ctx) error {
	log.Println("->", c.Method(), c.Path())
	return c.Next()
})

api := app.Group("/api", requireAPIKey) // only for /api/*
api.Get("/data", func(c *fiber.Ctx) error { return c.SendString("data") })
// no key → 401 ; X-API-Key: secret → 200 data
```

## Built-in middleware worth knowing

| Package | Does |
| --- | --- |
| `middleware/logger` | Access logs |
| `middleware/recover` | Recover from panics |
| `middleware/cors` | CORS headers |
| `middleware/limiter` | Rate limiting |
| `middleware/requestid` | Request ID header |

:::caution[Pitfalls]
- Order matters: `app.Use` calls run in registration order, and must come before the routes they protect.
:::
