package standardlibrary

import (
	"html/template"
	"os"
)

func Example_htmlTemplate() {
	t := template.Must(template.New("p").Parse(`<p>{{.}}</p>`))
	t.Execute(os.Stdout, "<script>alert(1)</script>") // escaped automatically
	// Output: <p>&lt;script&gt;alert(1)&lt;/script&gt;</p>
}
