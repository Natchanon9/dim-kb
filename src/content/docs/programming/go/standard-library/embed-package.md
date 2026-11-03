---
title: "embed Package"
description: "Use embed.FS with templates and file servers"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`embed.FS` is an `fs.FS`, so anything that accepts an `fs.FS` can read embedded files: templates, HTTP file servers, migration tools.

## Templates from the binary

```go title="templates.go"
//go:embed templates/*.tmpl
var templateFS embed.FS

t := template.Must(template.ParseFS(templateFS, "templates/*.tmpl"))
t.ExecuteTemplate(os.Stdout, "hello.tmpl", "embed") // Hello, embed!
```

## Static files

```go
//go:embed static
var static embed.FS

mux.Handle("GET /static/", http.FileServerFS(static)) // Go 1.22+
```

Migration tools such as golang-migrate and goose can also read SQL files from an `embed.FS`.

The directive syntax and rules are on [embed Directive](/dim-kb/programming/go/file-io/embed-directive/).

:::caution[Pitfalls]
- Embedded files are read-only and fixed at build time; config that changes per environment belongs in env vars.
:::
