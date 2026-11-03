---
title: "Break และ Continue"
description: "ออกจาก loop ก่อนกำหนด หรือข้ามไปรอบถัดไป"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`break` ออกจาก `for`, `switch` หรือ `select` ที่อยู่ในสุด ส่วน `continue` ข้ามไปรอบถัดไปของ loop ที่อยู่ในสุด

## ตัวอย่าง

```go title="skip.go"
for i := range 10 {
	if i%2 == 0 {
		continue // ข้ามเลขคู่
	}
	if i > 7 {
		break // หยุด loop
	}
	fmt.Print(i, " ")
}
// 1 3 5 7
```

:::caution[ข้อควรระวัง]
- ใน `for { select { ... } }` หรือ `for { switch { ... } }` คำสั่ง `break` เฉยๆ ออกแค่ `select`/`switch` แต่ loop ยังวนต่อ ให้ใช้ break แบบมี label หรือ `return`
:::
