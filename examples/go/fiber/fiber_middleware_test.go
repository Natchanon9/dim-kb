package gofiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func requireAPIKey(c *fiber.Ctx) error {
	if c.Get("X-API-Key") != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.Next() // continue to the next handler
}

func Example_middleware() {
	app := fiber.New()
	app.Use(recover.New())             // built-in: turns panics into 500s
	app.Use(func(c *fiber.Ctx) error { // custom logger
		fmt.Println("->", c.Method(), c.Path())
		return c.Next()
	})

	api := app.Group("/api", requireAPIKey) // middleware for a route group
	api.Get("/data", func(c *fiber.Ctx) error { return c.SendString("data") })

	call(app, "GET", "/api/data", "")
	call(app, "GET", "/api/data", "", "X-API-Key", "secret")
	// Output:
	// -> GET /api/data
	// 401 {"error":"unauthorized"}
	// -> GET /api/data
	// 200 data
}
