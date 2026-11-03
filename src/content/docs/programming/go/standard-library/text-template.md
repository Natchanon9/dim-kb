---
title: "text/template"
description: "Generate text from templates with data"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`text/template` fills `{{ }}` actions with data. Use it for emails, config files, code generation and CLI output.

## Example

```go title="template.go"
type order struct {
	Name  string
	VIP   bool
	Items []string
}

tmpl := template.Must(template.New("msg").Parse(
	"Hi {{.Name}}{{if .VIP}} (VIP){{end}}:{{range .Items}} {{.}};{{end}}\n"))
tmpl.Execute(os.Stdout, order{Name: "Ann", VIP: true, Items: []string{"pen", "ink"}})
// Hi Ann (VIP): pen; ink;
```

| Action | Meaning |
| --- | --- |
| `{{.Field}}` | Value from data |
| `{{if .X}}…{{else}}…{{end}}` | Condition |
| `{{range .Items}}{{.}}{{end}}` | Loop |
| `{{template "name" .}}` | Include another template |
| `{{.Name \| printf "%q"}}` | Pipe to a function |
| `{{- ... -}}` | Trim surrounding whitespace |

Add your own helpers with `.Funcs(template.FuncMap{...})` before `Parse`.

:::caution[Pitfalls]
- Never use `text/template` for HTML: it doesn't escape, so user input becomes XSS.
:::
