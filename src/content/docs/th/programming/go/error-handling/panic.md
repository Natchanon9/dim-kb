---
title: "panic"
description: "หยุดโปรแกรมเมื่อเจอ bug หรือสถานะที่เป็นไปไม่ได้"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`panic` หยุดการทำงานปกติ รัน deferred call ไล่ขึ้นไปตาม stack แล้วโปรแกรมพังพร้อม stack trace เว้นแต่มีจุดที่ recover

## ตัวอย่าง

```go title="panic.go"
func mustPositive(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("mustPositive: got %d", n))
	}
	return n
}
```

runtime ก็ panic เองได้: nil dereference, index เกินขอบ, เขียนลง nil map, ปิด channel ที่ปิดไปแล้ว

## เมื่อไหร่ควร panic

- ความผิดพลาดของโปรแกรมเมอร์ที่ไม่ควรเกิดตอนรันจริง
- ตอนเริ่มโปรแกรมที่ทำงานต่อไม่ได้แล้ว (`regexp.MustCompile` กับ pattern คงที่, ไม่มี config ที่จำเป็นใน `main`)

อะไรที่ผู้เรียกน่าจะจัดการได้ (input ผิด, network, ไฟล์) ให้เป็น `error` ไม่ใช่ panic

:::caution[ข้อควรระวัง]
- panic ใน goroutine ไหนก็ตามจะทำให้ทั้งโปรแกรมพัง ไม่ใช่แค่ goroutine นั้น
- library ไม่ควร panic ออกมาทาง API
:::
