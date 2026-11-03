---
title: "Nil Channels"
description: "operation บน nil channel บล็อกตลอดไป ใช้ปิด case ใน select ได้"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

ส่งหรือรับจาก channel ที่เป็น `nil` จะบล็อกตลอดไป ส่วนการปิดจะ panic ใน `select` การบล็อกนี้มีประโยชน์: ตั้ง channel เป็น nil เพื่อปิด case นั้น

## รวมสอง stream จนกว่าจะปิดทั้งคู่

```go title="nilcase.go"
a, b := emit(1), emit(2)
sum := 0
for a != nil || b != nil {
	select {
	case v, ok := <-a:
		if !ok {
			a = nil // case นี้จะไม่ทำงานอีก
			continue
		}
		sum += v
	case v, ok := <-b:
		if !ok {
			b = nil
			continue
		}
		sum += v
	}
}
fmt.Println(sum) // 3
```

ถ้าไม่ใช้ nil channel ที่ปิดแล้วจะ "พร้อม" ตลอดเวลา และ loop จะหมุนไม่หยุด

:::caution[ข้อควรระวัง]
- ลืม `make` ทำให้ channel เป็น nil goroutine ที่ใช้มันจะค้างเงียบๆ
:::
