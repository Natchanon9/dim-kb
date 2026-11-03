---
title: "html/template"
description: "template HTML ที่ escape อัตโนมัติตามบริบท"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`html/template` ใช้ syntax เดียวกับ `text/template` แต่ escape ข้อมูลตามตำแหน่งที่วาง (HTML, attribute, URL, JavaScript, CSS) จึงกัน XSS ได้ตั้งแต่ต้น

## ตัวอย่าง

```go title="page.go"
t := template.Must(template.New("p").Parse(`<p>{{.}}</p>`))
t.Execute(os.Stdout, "<script>alert(1)</script>")
// <p>&lt;script&gt;alert(1)&lt;/script&gt;</p>
```

## Render ฝั่ง server

```go
//go:embed views/*.html
var views embed.FS
tmpl := template.Must(template.ParseFS(views, "views/*.html"))

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", data)
}
```

parse template ครั้งเดียวตอนเริ่มโปรแกรม ไม่ใช่ทุก request

:::caution[ข้อควรระวัง]
- `template.HTML(s)` บอกว่า string นั้นปลอดภัยและปิดการ escape ใช้กับเนื้อหาที่เราควบคุมเองเท่านั้น
:::
