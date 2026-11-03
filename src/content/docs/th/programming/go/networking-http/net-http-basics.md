---
title: "พื้นฐาน net/http"
description: "HTTP server และ client ระดับ production ใน standard library"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`net/http` มีทั้ง server และ client ของ HTTP/1.1 และ HTTP/2 ครบ service Go หลายตัวใช้ตรงๆ โดยไม่มี framework

## Server ที่เล็กที่สุด

```go title="main.go"
mux := http.NewServeMux()
mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
})
log.Fatal(http.ListenAndServe(":8080", mux))
// GET /hello → 200 hello
```

## ส่วนประกอบ

| ส่วน | หน้าที่ |
| --- | --- |
| `http.Handler` | อะไรก็ได้ที่มี `ServeHTTP(w, r)` |
| `http.HandlerFunc` | แปลง function เป็น Handler |
| `http.ServeMux` | router: pattern → handler |
| `http.Server` | listener ที่มี timeout และ shutdown |
| `http.Client` | ส่ง request ออกไปข้างนอก |

แต่ละ request รันใน goroutine ของตัวเอง handler จึงต้องปลอดภัยเมื่อถูกเรียกพร้อมกัน

:::caution[ข้อควรระวัง]
- `http.ListenAndServe` ไม่มี timeout ใน production ให้สร้าง `http.Server` ที่มี `ReadHeaderTimeout` (และส่วนใหญ่ก็ `ReadTimeout`, `WriteTimeout`, `IdleTimeout`)
:::
