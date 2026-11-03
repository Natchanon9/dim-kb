---
title: "พื้นฐาน cgo"
description: "เรียกโค้ด C จาก Go ด้วย import \"C\""
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

cgo ให้ Go เรียก C ได้ โค้ด C อยู่ใน comment ที่อยู่เหนือ `import "C"` พอดี แล้ว function ของ C จะเรียกผ่าน prefix `C.`

## ตัวอย่าง

```go title="cgo.go"
//go:build cgo

package mathc

/*
static int add(int a, int b) { return a + b; }
*/
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
// Add(2, 3) → 5
```

## สิ่งที่ต้องมีตอน build

- C compiler (gcc/clang) บนเครื่องที่ build
- `CGO_ENABLED=1` (เปิดให้อัตโนมัติถ้าเจอ compiler แต่ปิดตอน cross-compile)
- tag `//go:build cgo` ทำให้ package build ได้แม้ไม่มี cgo ถ้าเรามีทางสำรองให้

:::caution[ข้อควรระวัง]
- `import "C"` ต้องตามหลัง comment ของ C ทันที ห้ามมีบรรทัดว่างคั่น
- ใช้ cgo ในไฟล์ `_test.go` ไม่ได้ ต้องวางในไฟล์ปกติแล้ว test ตัวนั้น
:::
