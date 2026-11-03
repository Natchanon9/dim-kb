---
title: "คืนค่าหลายค่า"
description: "คืนผลลัพธ์กับ error พร้อมกัน"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

function ของ Go คืนค่าได้หลายค่า รูปแบบมาตรฐานคือ `(result, error)` โดย error อยู่ท้ายสุด

## รูปแบบ `(value, error)`

```go title="divide.go"
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

q, err := divide(10, 4) // 2.5 <nil>
_, err = divide(1, 0)   // divide by zero
```

เช็ค `err` ก่อนเสมอ ถ้าไม่เป็น nil ให้ถือว่าค่าอื่นใช้ไม่ได้

:::caution[ข้อควรระวัง]
- ต้องรับผลลัพธ์ครบทุกตัวหรือไม่รับเลย: `x := divide(1, 2)` compile ไม่ผ่าน ใช้ `_` ทิ้งตัวที่ไม่ต้องการ
- คืนค่าเกิน 2–3 ตัวแปลว่าควรใช้ struct แทน
:::
