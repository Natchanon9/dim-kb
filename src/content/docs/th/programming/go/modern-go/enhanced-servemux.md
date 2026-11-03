---
title: "Enhanced ServeMux"
description: "pattern ที่มี method, wildcard และลำดับความสำคัญใน Go 1.22"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

Go 1.22 อัปเกรด `http.ServeMux` ให้มี method, `{wildcard}` และลำดับความสำคัญที่ชัดเจน ตัดเหตุผลหลักที่ต้องเพิ่ม router library ออกไป

## ตัวอย่าง

```go title="routes.go"
mux.HandleFunc("GET /{$}", home)              // "/" เท่านั้น
mux.HandleFunc("GET /posts/{id}", showPost)   // r.PathValue("id")
mux.HandleFunc("GET /posts/latest", latest)   // เจาะจงกว่า: ชนะ {id}

// /             → 200 home
// /posts/7      → 200 post 7
// /posts/latest → 200 latest
// /nope         → 404
```

## อะไรเปลี่ยนจาก 1.21

| ก่อน | ตั้งแต่ 1.22 |
| --- | --- |
| path อย่างเดียว | `METHOD /path` |
| ไม่มี parameter | `{name}`, `{rest...}` |
| prefix ยาวสุดชนะ | pattern เจาะจงสุดชนะ |
| `/` จับทุกอย่าง | `/{$}` สำหรับ root พอดี |

ใน test ใช้ `r.SetPathValue` เพื่อจำลอง wildcard ตอนเรียก handler ตรงๆ รายละเอียดอยู่ที่ [ServeMux Routing](/dim-kb/th/programming/go/networking-http/servemux-routing/)

:::caution[ข้อควรระวัง]
- พฤติกรรมเก่ากลับมาถ้า `go.mod` เป็น `go 1.21` หรือต่ำกว่า หรือตั้ง `GODEBUG=httpmuxgo121=1`
:::
