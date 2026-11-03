---
title: "errors.Is และ errors.As"
description: "เช็คว่ามี error ค่าหรือ type ที่ต้องการอยู่ตรงไหนในสายหรือไม่"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`errors.Is` ถามว่า "มี error ตัวนี้อยู่ในสายไหม" ส่วน `errors.As` ถามว่า "มี error ที่เป็น type นี้ไหม ถ้ามีขอตัวนั้นมา"

## ตัวอย่าง

```go title="isas.go"
var ErrNotFound = errors.New("not found")

err := fmt.Errorf("get user: %w", ErrNotFound)
errors.Is(err, ErrNotFound) // true

err = fmt.Errorf("save: %w", validateEmail("bob"))
var ve *ValidationError
if errors.As(err, &ve) {
	fmt.Println(ve.Field) // email
}
```

| Function | เทียบ | เป้าหมายทั่วไป |
| --- | --- | --- |
| `errors.Is(err, target)` | ค่า | sentinel: `io.EOF`, `sql.ErrNoRows` |
| `errors.As(err, &target)` | type | custom type: `*ValidationError`, `*fs.PathError` |

:::caution[ข้อควรระวัง]
- `errors.As` ต้องรับ **pointer** ไปยังตัวแปรของ type เป้าหมาย: `&ve` โดยที่ `ve` เป็น `*ValidationError`
- `err == ErrNotFound` ใช้ไม่ได้ทันทีที่ error ถูก wrap ให้ใช้ `errors.Is` เสมอ
:::
