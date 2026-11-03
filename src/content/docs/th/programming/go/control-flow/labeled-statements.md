---
title: "Labeled Statements"
description: "break และ continue loop ด้านนอกด้วย label"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

label คือการตั้งชื่อให้ loop เพื่อให้ `break` หรือ `continue` จาก loop ซ้อนหรือ `select` ข้างในเล็งไปที่ loop นั้นได้

## ตัวอย่าง

```go title="labels.go"
outer:
	for i := range 3 {
		for j := range 3 {
			if j == 2 {
				continue outer // ไป i ถัดไป
			}
			if i == 2 {
				break outer // ออกทั้งสอง loop
			}
			fmt.Printf("%d%d ", i, j)
		}
	}
// 00 01 10 11
```

การใช้จริงที่เจอบ่อยสุดคือออกจาก loop `for { select { ... } }` ตอน channel ปิด

:::caution[ข้อควรระวัง]
- label ที่ไม่ได้ใช้ compile ไม่ผ่าน
- ถ้าต้องใช้ label บ่อย ให้แยก loop ด้านในเป็น function แล้วใช้ `return` มักอ่านง่ายกว่า
:::
