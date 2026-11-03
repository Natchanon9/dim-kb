---
title: "embed Package"
description: "ใช้ embed.FS กับ template และ file server"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`embed.FS` เป็น `fs.FS` อะไรที่รับ `fs.FS` จึงอ่านไฟล์ที่ embed ได้หมด: template, HTTP file server, เครื่องมือ migration

## Template จากใน binary

```go title="templates.go"
//go:embed templates/*.tmpl
var templateFS embed.FS

t := template.Must(template.ParseFS(templateFS, "templates/*.tmpl"))
t.ExecuteTemplate(os.Stdout, "hello.tmpl", "embed") // Hello, embed!
```

## Static file

```go
//go:embed static
var static embed.FS

mux.Handle("GET /static/", http.FileServerFS(static)) // Go 1.22+
```

เครื่องมือ migration อย่าง golang-migrate และ goose ก็อ่านไฟล์ SQL จาก `embed.FS` ได้

syntax และกติกาของ directive อยู่ที่ [embed Directive](/dim-kb/th/programming/go/file-io/embed-directive/)

:::caution[ข้อควรระวัง]
- ไฟล์ที่ embed อ่านได้อย่างเดียวและตายตัวตั้งแต่ตอน build config ที่เปลี่ยนตาม environment ให้ใช้ env var
:::
