---
title: "html/template"
description: "HTML templates with automatic, context-aware escaping"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`html/template` has the same syntax as `text/template` but escapes data based on where it appears (HTML, attribute, URL, JavaScript, CSS). This blocks XSS by default.

## Example

```go title="page.go"
t := template.Must(template.New("p").Parse(`<p>{{.}}</p>`))
t.Execute(os.Stdout, "<script>alert(1)</script>")
// <p>&lt;script&gt;alert(1)&lt;/script&gt;</p>
```

## Server-side rendering

```go
//go:embed views/*.html
var views embed.FS
tmpl := template.Must(template.ParseFS(views, "views/*.html"))

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", data)
}
```

Parse templates once at startup, not per request.

:::caution[Pitfalls]
- `template.HTML(s)` marks a string as safe and turns escaping off; use it only for content you control.
:::
