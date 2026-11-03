---
title: "cgo Pitfalls"
description: "ต้นทุนการเรียก การ build การ deploy และความปลอดภัย"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

"cgo ไม่ใช่ Go" ใช้ได้ แต่ต้องยอมเสียหลายอย่างที่ทำให้ Go ใช้สบาย

## ต้นทุนของการเรียก

```go title="bench_test.go"
func BenchmarkGoAdd(b *testing.B)  { /* บวกด้วย Go ธรรมดา */ }
func BenchmarkCgoAdd(b *testing.B) { /* Add ผ่าน C */ }
// GoAdd   ~0.2 ns/op
// CgoAdd  ~18 ns/op  (ทุกครั้งต้องข้ามขอบ Go/C)
```

รวบงานเป็นก้อนเพื่อข้ามขอบให้น้อยครั้ง

## ต้นทุนอื่น

| ด้าน | ปัญหา |
| --- | --- |
| Build | ต้องมี C toolchain และ compile ช้าลง |
| Cross-compile | `GOOS=linux go build` ใช้ตรงๆ ไม่ได้แล้ว |
| Deploy | binary อาจพึ่ง shared lib (ใช้ image `scratch` ไม่ได้) |
| ความปลอดภัย | bug ใน C ทำทั้ง process พัง และ race detector มองไม่เห็นใน C |
| Tooling | profiler และ debugger เห็นน้อยลง |

## ทางเลือก

ใช้ library ที่เป็น Go ล้วน (เช่น `modernc.org/sqlite` แทน `mattn/go-sqlite3`), เรียกเป็น subprocess หรือห่อโค้ด C เป็น service แยกบน network

:::caution[ข้อควรระวัง]
- การเรียก C ที่บล็อกจะกิน OS thread ตัวละหนึ่ง ถ้าเรียกพร้อมกันเป็นพันก็จะได้ thread เป็นพัน
:::
