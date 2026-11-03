package gofiber

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func newBookApp() *fiber.App {
	app := fiber.New()
	var (
		mu     sync.Mutex
		books  = map[int]Book{}
		nextID = 1
	)

	app.Post("/books", func(c *fiber.Ctx) error {
		var b Book
		if err := c.BodyParser(&b); err != nil || b.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title required"})
		}
		mu.Lock()
		b.ID = nextID
		nextID++
		books[b.ID] = b
		mu.Unlock()
		return c.Status(fiber.StatusCreated).JSON(b)
	})

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		mu.Lock()
		b, ok := books[id]
		mu.Unlock()
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.JSON(b)
	})

	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		mu.Lock()
		delete(books, id)
		mu.Unlock()
		return c.SendStatus(fiber.StatusNoContent)
	})
	return app
}

func Example_crud() {
	app := newBookApp()
	call(app, "POST", "/books", `{"title":"Go"}`)
	call(app, "POST", "/books", `{}`)
	call(app, "GET", "/books/1", "")
	call(app, "DELETE", "/books/1", "")
	call(app, "GET", "/books/1", "")
	// Output:
	// 201 {"id":1,"title":"Go"}
	// 400 {"error":"title required"}
	// 200 {"id":1,"title":"Go"}
	// 204
	// 404 {"error":"not found"}
}
