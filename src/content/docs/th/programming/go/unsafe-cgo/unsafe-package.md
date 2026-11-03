---
title: "unsafe Package"
description: "ออกนอกระบบความปลอดภัยของ type และทำไมแทบไม่ควรทำ"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`unsafe` ให้เข้าถึงรูปแบบ memory และแปลง pointer ระดับล่าง โค้ดที่ import มันจะเสียการรับประกันเรื่องความปลอดภัยของ memory และ compatibility ของ Go

## ข้างในมีอะไร

```go title="sizes.go"
fmt.Println(unsafe.Sizeof(int64(0)), unsafe.Sizeof(""), unsafe.Sizeof([]int{}))
// 8 16 24  (64-bit: string = ptr+len, slice = ptr+len+cap)
```

| สิ่งที่มี | ทำอะไร |
| --- | --- |
| `unsafe.Pointer` | แปลงระหว่าง pointer type ใดก็ได้ |
| `Sizeof`, `Alignof`, `Offsetof` | รูปแบบ memory ตอน compile |
| `String`, `StringData`, `Slice`, `SliceData` (1.20+) | มอง string/slice แบบไม่ copy |
| `Add` | pointer arithmetic |

## เมื่อไหร่ถึงสมเหตุสมผล

library ที่ต้องการ performance สูง (serialization, hashing), syscall และการทำงานร่วมกับ C โค้ดแอปแทบไม่เคยต้องใช้

:::caution[ข้อควรระวัง]
- โค้ดที่ใช้ `unsafe` อาจพังเมื่อเปลี่ยนเวอร์ชัน Go หรือเปลี่ยน architecture เพราะ Go 1 compatibility promise ไม่ครอบคลุม
:::
