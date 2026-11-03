---
title: "embed Directive"
description: "ฝังไฟล์ลงใน binary ด้วย //go:embed"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`//go:embed` (Go 1.16+) ใส่ไฟล์ลงใน binary ที่ compile แล้ว template, SQL migration และ static asset จึงไปพร้อมกับ executable ไฟล์เดียว

## ตัวอย่าง

```go title="assets.go"
import "embed"

//go:embed static/hello.txt
var hello string // เนื้อหาไฟล์เป็น string ([]byte ก็ได้)

//go:embed static
var staticFS embed.FS // ทั้งโฟลเดอร์เป็น fs.FS

fmt.Print(hello) // hello from embed
entries, _ := fs.ReadDir(staticFS, "static")
```

serve asset ที่ embed ได้ตรงๆ: `http.FileServerFS(staticFS)` (Go 1.22+)

:::caution[ข้อควรระวัง]
- directive ต้องอยู่เหนือ `var` ระดับ package พอดี และ package ต้อง import `embed`
- path สัมพัทธ์กับไฟล์ source และใช้ `..` ไม่ได้ ไฟล์ที่ขึ้นต้นด้วย `.` หรือ `_` จะถูกข้าม เว้นแต่ใช้ `all:`
:::
