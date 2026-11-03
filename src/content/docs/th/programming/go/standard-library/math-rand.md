---
title: "math/rand"
description: "เลขสุ่มเทียมด้วย math/rand/v2"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

`math/rand/v2` (Go 1.22+) สร้างเลขสุ่มเทียมที่เร็ว สำหรับ simulation, jitter, การสุ่มตัวอย่าง และการสลับลำดับ function ระดับบนสุดตั้ง seed ให้อัตโนมัติ

## ตัวอย่าง

```go title="rand.go"
rand.IntN(10)           // 0..9
rand.Float64()          // [0.0, 1.0)
rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
rand.N(5 * time.Second) // ระยะเวลาสุ่ม เช่น jitter ของการ retry

a := rand.New(rand.NewPCG(1, 2)) // seed ตายตัว: ได้ลำดับเดิมทุกครั้ง
b := rand.New(rand.NewPCG(1, 2))
a.IntN(100) == b.IntN(100) // true
```

## v1 vs v2

| `math/rand` (v1) | `math/rand/v2` |
| --- | --- |
| `Intn` | `IntN` |
| ก่อน Go 1.20 ต้อง `rand.Seed` เอง | ตั้ง seed ให้อัตโนมัติเสมอ |
| source แบบเดียว | PCG, ChaCha8 |

:::caution[ข้อควรระวัง]
- ไม่ใช่สำหรับรหัสผ่าน token หรือ key ให้ใช้ `crypto/rand`
- `*rand.Rand` ที่มี source ของตัวเองใช้พร้อมกันหลาย goroutine ไม่ปลอดภัย แต่ function ระดับบนสุดปลอดภัย
:::
