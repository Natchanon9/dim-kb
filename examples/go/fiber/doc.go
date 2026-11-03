// Package gofiber holds the snippets for the "Fiber" pages (Backend Practice)
// in src/content/docs/programming/go/fiber/.
package gofiber

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// call sends one request through app without opening a port and prints status and body.
func call(app *fiber.App, method, path, body string, headers ...string) {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	for i := 0; i+1 < len(headers); i += 2 {
		req.Header.Set(headers[i], headers[i+1])
	}
	res, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(res.Body)
	fmt.Println(strings.TrimSpace(fmt.Sprint(res.StatusCode, " ", strings.TrimSpace(string(b)))))
}
