---
title: "Loop Variable Change"
description: "ตั้งแต่ Go 1.22 แต่ละรอบของ loop ได้ตัวแปรของตัวเอง"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

ก่อน Go 1.22 loop `for` มีตัวแปรตัวเดียวที่ทุกรอบใช้ร่วมกัน closure และ goroutine จึงจับได้ค่าสุดท้าย ตั้งแต่ 1.22 แต่ละรอบมีสำเนาของตัวเอง

## ตัวอย่าง

```go title="loopvar.go"
var fns []func() int
for i := range 3 {
	fns = append(fns, func() int { return i })
}
for _, f := range fns {
	fmt.Print(f()) // 012 บน Go 1.22+ (เดิม loop แบบ 3 ส่วนได้ 333)
}
```

## แก้ปัญหาอะไร

- `go func() { use(v) }()` ใน `range` เคยเห็นแต่ element ตัวสุดท้าย
- เอา `&v` ใน loop เคยได้ pointer ตัวเดียวกันทุกรอบ

วิธีแก้แบบเก่า `v := v` ไม่จำเป็นแล้ว และเครื่องมืออย่าง `copyloopvar` จะทักให้ลบ

:::caution[ข้อควรระวัง]
- พฤติกรรมใหม่ขึ้นกับบรรทัด `go` ใน `go.mod` ราย module module เก่ายังใช้พฤติกรรมเดิม
:::
