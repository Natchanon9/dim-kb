---
title: "sync.Once"
description: "รัน initialization ครั้งเดียวอย่างปลอดภัย"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`sync.Once` รัน function ครั้งเดียวเท่านั้น แม้หลาย goroutine จะเรียกพร้อมกัน ตัวอื่นจะรอจนการเรียกครั้งแรกเสร็จ

## ตัวอย่าง

```go title="once.go"
var once sync.Once
calls := 0
for range 3 {
	once.Do(func() { calls++ }) // รันแค่ครั้งแรก
}

config := sync.OnceValue(func() string { // Go 1.21+
	return "loaded"
})
fmt.Println(calls, config(), config()) // 1 loaded loaded
```

`OnceValue` กับ `OnceValues` (ค่า + error) cache ผลลัพธ์ไว้ให้ ตรงกับความต้องการแบบ lazy singleton ทั่วไป

:::caution[ข้อควรระวัง]
- ถ้า function panic `Once` ก็ยังนับว่าทำแล้ว เรียกครั้งต่อไปจะไม่ทำอะไร
- เรียก `once.Do` ข้างใน function ที่มันรันอยู่จะ deadlock
:::
