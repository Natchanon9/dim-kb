---
title: "พื้นฐาน Context"
description: "ส่งต่อการยกเลิก deadline และค่าของ request ข้ามการเรียก function"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`context.Context` พกสัญญาณยกเลิก deadline (ถ้ามี) และค่าที่ผูกกับ request ทุก function ที่บล็อกใน server ควรรับมัน

## ตัวอย่าง

```go title="ctx.go"
func slowOp(ctx context.Context) error {
	select {
	case <-time.After(time.Second): // งานจริง
		return nil
	case <-ctx.Done(): // ถูกยกเลิกหรือหมดเวลา
		return ctx.Err()
	}
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
defer cancel()
fmt.Println(slowOp(ctx)) // context deadline exceeded
```

## Interface

| Method | คืนค่า |
| --- | --- |
| `Done()` | channel ที่ปิดเมื่อ context จบ |
| `Err()` | `nil`, `context.Canceled` หรือ `context.DeadlineExceeded` |
| `Deadline()` | deadline (ถ้ามี) |
| `Value(key)` | ค่าที่ผูกกับ request |

จุดเริ่ม: `context.Background()` ใน `main`/test และ `context.TODO()` ใช้เป็นตัวแทนชั่วคราว ใน HTTP handler ใช้ `r.Context()`

:::caution[ข้อควรระวัง]
- ส่ง `ctx` เป็น parameter **ตัวแรก** และอย่าเก็บไว้ใน struct
- อย่าส่ง context ที่เป็น `nil` ให้ใช้ `context.TODO()`
:::
