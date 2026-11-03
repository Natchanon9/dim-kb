---
title: "HTTP Middleware"
description: "ห่อ handler เพื่อเพิ่ม logging, auth และ recovery"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

middleware คือ function ที่รับ handler แล้วคืน handler ตัวใหม่ รันโค้ดก่อนและหลัง handler ที่ห่ออยู่ หรือหยุด request ก่อนก็ได้

## ตัวอย่าง

```go title="middleware.go"
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("->", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func requireKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != "secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return // หยุดสาย
		}
		next.ServeHTTP(w, r)
	})
}

h := logging(requireKey(dataHandler)) // ชั้นนอกสุดรันก่อน
// GET /data ไม่มี key → "-> GET /data", 401
```

middleware ที่เจอบ่อย: logging, request ID, auth, CORS, panic recovery, rate limiting, timeout (`http.TimeoutHandler`)

:::caution[ข้อควรระวัง]
- ถ้าจะ log status code ต้องห่อ `ResponseWriter` ด้วย struct ที่จำค่าจาก `WriteHeader` ไว้
:::
