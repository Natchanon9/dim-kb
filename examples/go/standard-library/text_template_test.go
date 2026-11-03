package standardlibrary

import (
	"os"
	"text/template"
)

type order struct {
	Name  string
	VIP   bool
	Items []string
}

func Example_textTemplate() {
	tmpl := template.Must(template.New("msg").Parse(
		"Hi {{.Name}}{{if .VIP}} (VIP){{end}}:{{range .Items}} {{.}};{{end}}\n"))
	tmpl.Execute(os.Stdout, order{Name: "Ann", VIP: true, Items: []string{"pen", "ink"}})
	// Output: Hi Ann (VIP): pen; ink;
}
