package standardlibrary

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

func Example_embedPackage() {
	t := template.Must(template.ParseFS(templateFS, "templates/*.tmpl")) // parse straight from the binary
	t.ExecuteTemplate(os.Stdout, "hello.tmpl", "embed")
	// Output: Hello, embed!
}
