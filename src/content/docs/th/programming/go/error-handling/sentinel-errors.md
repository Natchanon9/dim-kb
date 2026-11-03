---
title: "Sentinel Errors"
description: "ค่า error ที่ export ไว้ให้ผู้เรียกเปรียบเทียบ"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

sentinel คือค่า error ระดับ package ที่ผู้เรียกเช็คด้วย `errors.Is` ตั้งชื่อเป็น `ErrXxx`

## ตัวอย่าง

```go title="stock.go"
var ErrOutOfStock = errors.New("out of stock")

func buy(stock int) error {
	if stock == 0 {
		return fmt.Errorf("buy: %w", ErrOutOfStock)
	}
	return nil
}

if err := buy(0); errors.Is(err, ErrOutOfStock) {
	fmt.Println("show sold-out badge")
}
```

ตัวมาตรฐานที่จะเจอ: `io.EOF`, `sql.ErrNoRows`, `context.Canceled`, `context.DeadlineExceeded`, `os.ErrNotExist`

## Sentinel หรือ custom type?

ใช้ sentinel เมื่อผู้เรียกแค่ต้องรู้ว่าเป็น error **ตัวไหน** ใช้ custom type เมื่อต้องการ **ข้อมูล** จาก error

:::caution[ข้อควรระวัง]
- sentinel เป็นส่วนหนึ่งของ public API เปลี่ยนชื่อหรือลบทิ้งจะทำให้ผู้เรียกพัง
:::
