---
title: "Swagger"
description: "Generate OpenAPI docs from comments with swag"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`swaggo/swag` reads annotation comments on handlers and generates an OpenAPI (Swagger) spec; `gofiber/swagger` serves Swagger UI.

## Annotate handlers

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

Add general info (`@title`, `@version`, `@BasePath`) above `main`.

## Generate and serve

```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init            # writes docs/ (swagger.json, swagger.yaml, docs.go)
```

```go
import (
	_ "myapp/docs"
	"github.com/gofiber/swagger"
)

app.Get("/swagger/*", swagger.HandlerDefault) // UI at /swagger/index.html
```

:::caution[Pitfalls]
- Annotations drift from code silently; run `swag init` in CI (or `go generate`) and fail on a diff.
:::
