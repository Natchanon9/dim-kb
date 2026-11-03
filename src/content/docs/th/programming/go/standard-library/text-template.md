---
title: "text/template"
description: "สร้างข้อความจาก template กับข้อมูล"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`text/template` เติมข้อมูลลงใน action `{{ }}` ใช้กับ email, ไฟล์ config, การ generate โค้ด และ output ของ CLI

## ตัวอย่าง

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

| Action | ความหมาย |
| --- | --- |
| `{{.Field}}` | ค่าจากข้อมูล |
| `{{if .X}}…{{else}}…{{end}}` | เงื่อนไข |
| `{{range .Items}}{{.}}{{end}}` | วน loop |
| `{{template "name" .}}` | ใส่ template อื่น |
| `{{.Name \| printf "%q"}}` | ส่งต่อให้ function |
| `{{- ... -}}` | ตัดช่องว่างรอบๆ |

เพิ่ม helper ของเราเองด้วย `.Funcs(template.FuncMap{...})` ก่อน `Parse`

:::caution[ข้อควรระวัง]
- ห้ามใช้ `text/template` กับ HTML เพราะไม่ escape ให้ input ของผู้ใช้จะกลายเป็น XSS
:::
