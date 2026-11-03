---
title: "recover"
description: "จับ panic ภายใน deferred function"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`recover()` หยุด panic และคืนค่าของมัน แต่ได้ผลเฉพาะเมื่อเรียกตรงๆ ใน deferred function

## แปลง panic เป็น error

```go title="recover.go"
func safeDivide(a, b int) (q int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

safeDivide(6, 3) // 2 <nil>
safeDivide(1, 0) // 0 recovered: runtime error: integer divide by zero
```

## ใช้จริงที่ไหน

HTTP server recover ทีละ request เพื่อไม่ให้ handler ตัวเดียวทำ server ล่ม (`net/http` ทำให้อยู่แล้ว ส่วน framework มี recovery middleware) goroutine ที่เป็น worker ก็มัก recover แล้วรายงานเช่นกัน

:::caution[ข้อควรระวัง]
- `recover` จับได้เฉพาะ panic ใน **goroutine เดียวกัน** ทุก goroutine ที่สร้างต้องมีของตัวเอง
- อย่าใช้ panic/recover เป็น try/catch กับ error ปกติ
:::
