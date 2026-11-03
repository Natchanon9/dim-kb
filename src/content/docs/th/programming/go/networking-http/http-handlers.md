---
title: "HTTP Handlers"
description: "Handler, HandlerFunc และ handler ที่มี dependency"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

handler ได้ `ResponseWriter` ไว้เขียน response และ `*Request` ไว้อ่าน input type ไหนมี `ServeHTTP(w, r)` ก็เป็น `http.Handler`

## Handler แบบ struct ที่มี dependency

```go title="greeter.go"
type greeter struct{ greeting string } // DB, logger, config อยู่ตรงนี้

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s", g.greeting, r.URL.Query().Get("name"))
}
// GET /?name=ann → hi, ann
```

## อ่าน request

| ต้องการ | ใช้ |
| --- | --- |
| path param | `r.PathValue("id")` (Go 1.22+) |
| query string | `r.URL.Query().Get("q")` |
| header | `r.Header.Get("Authorization")` |
| JSON body | `json.NewDecoder(r.Body).Decode(&v)` |
| form | `r.FormValue("email")` |
| การยกเลิก | `r.Context()` |

## เขียน response

ตั้ง header ก่อน แล้ว `w.WriteHeader(status)` แล้วค่อยเขียน body ถ้าเขียน body ก่อนจะส่ง `200` ให้อัตโนมัติ

:::caution[ข้อควรระวัง]
- ตั้ง header หลัง `WriteHeader` หรือหลัง `Write` ครั้งแรกไม่มีผล
- หลัง `http.Error(...)` ต้อง `return` ไม่อย่างนั้น handler จะเขียนต่อ
:::
