package gofiber

import "github.com/gofiber/fiber/v2"

// GetBook godoc
//
//	@Summary	Get a book by ID
//	@Tags		books
//	@Produce	json
//	@Param		id	path		int	true	"Book ID"
//	@Success	200	{object}	Book
//	@Failure	404	{object}	map[string]string
//	@Router		/books/{id} [get]
func GetBook(c *fiber.Ctx) error {
	return c.JSON(Book{ID: 1, Title: "Go"})
}

func Example_swaggerHandler() {
	app := fiber.New()
	app.Get("/books/:id", GetBook)
	call(app, "GET", "/books/1", "")
	// Output: 200 {"id":1,"title":"Go"}
}
