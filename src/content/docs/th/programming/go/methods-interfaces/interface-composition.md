---
title: "Interface Composition"
description: "สร้าง interface ใหญ่จากการ embed interface เล็ก"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

interface embed interface อื่นได้ ผลลัพธ์ต้องมี method ของทุกตัว `io.ReadWriter` ก็สร้างจาก `io.Reader` กับ `io.Writer` แบบนี้

## ตัวอย่าง

```go title="compose.go"
type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface {
	Reader
	Writer
}

var rw ReadWriter = &memFile{} // memFile มีทั้ง Read และ Write
rw.Write("hi")
fmt.Println(rw.Read()) // hi
```

## หลักการออกแบบ

เริ่มจาก interface ที่มี 1–2 method แล้วประกอบกันตรงที่ function ต้องการมากกว่านั้น function ควรขอ interface ที่เล็กที่สุดที่ใช้จริง

:::caution[ข้อควรระวัง]
- interface ใหญ่ (10+ method) implement ยาก และทำตัวปลอมตอน test ยาก ให้แยกย่อย
:::
