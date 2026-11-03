---
title: "unsafe.Pointer"
description: "แปลง pointer ข้าม type และการแปลงแบบไม่ copy"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`unsafe.Pointer` แปลงไปและกลับจาก pointer type ใดก็ได้ ข้ามระบบ type ไปเลย spec ระบุรูปแบบที่ถูกต้องไว้ไม่กี่แบบ

## ตัวอย่าง

```go title="pointer.go"
b := []byte("hello")
s := unsafe.String(unsafe.SliceData(b), len(b)) // ไม่ copy; ห้ามแก้ b อีกเลย
fmt.Println(s) // hello

x := int32(7)
p := unsafe.Pointer(&x)
fmt.Println(*(*int32)(p)) // 7
```

## กติกาที่ทำให้ยังถูกต้อง

- แปลง `*T1 → unsafe.Pointer → *T2` เฉพาะเมื่อรูปแบบ memory เข้ากันได้
- อย่าเก็บ pointer เป็น `uintptr` ข้าม statement เพราะ GC มองไม่เห็น object อาจย้ายที่หรือถูกคืนไปแล้ว
- เทคนิคแปลงเป็น string แบบไม่ copy ปลอดภัยก็ต่อเมื่อ byte นั้นไม่ถูกแก้อีก

`go vet` ตรวจการใช้ผิดบางแบบได้ ("possible misuse of unsafe.Pointer")

:::caution[ข้อควรระวัง]
- bug ตรงนี้โผล่เป็น crash แบบสุ่มหรือข้อมูลเสียในที่ที่ห่างไกลจากต้นเหตุ
:::
