---
title: "ServeMux Routing"
description: "pattern ที่มี method และ path parameter ใน Go 1.22+"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

ตั้งแต่ Go 1.22 pattern ของ `http.ServeMux` ใส่ HTTP method และ wildcard `{name}` ได้ ครอบคลุมงาน routing ส่วนใหญ่โดยไม่ต้องใช้ router ภายนอก

## ตัวอย่าง

```go title="routes.go"
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user ", r.PathValue("id"))
})
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /files/{path...}", serveFile) // {x...} จับส่วนที่เหลือทั้งหมด

// GET /users/42      → 200 user 42
// POST /users        → 201 created
// GET /files/a/b.txt → 200 file a/b.txt
// DELETE /users/42   → 405 Method Not Allowed
```

## กติกาการจับคู่

- pattern ที่ **เจาะจงที่สุด** ชนะ ไม่ขึ้นกับลำดับการลงทะเบียน
- `GET` จับ `HEAD` ด้วย
- `/` ท้าย (`/static/`) จับทั้ง subtree ส่วน `{$}` (`/{$}`) จับ path นั้นพอดีเท่านั้น
- pattern ที่ขัดกันจะ panic ตอนลงทะเบียน

:::caution[ข้อควรระวัง]
- `go.mod` ต้องเป็น `go 1.22` ขึ้นไป ไม่อย่างนั้นจะใช้กติกาเก่า และ `GET /x` จะถูกมองเป็น path ตรงตัว
:::
