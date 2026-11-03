---
title: "Closures"
description: "function ที่จับตัวแปรข้างนอกไว้และเก็บค่าต่อ"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

closure คือ function ที่ใช้ตัวแปรจาก scope ที่มันถูกสร้าง ตัวแปรพวกนั้นจะอยู่ตราบที่ closure ยังอยู่

## ตัวนับ

```go title="counter.go"
func counter() func() int {
	n := 0
	return func() int {
		n++ // n อยู่ต่อระหว่างการเรียกแต่ละครั้ง
		return n
	}
}

next := counter()
next(); next()
other := counter()           // มี n ของตัวเอง
fmt.Println(next(), other()) // 3 1
```

closure จับตัวแปรแบบ **by reference** การเรียก `counter()` แต่ละครั้งจึงได้ state แยกกัน

## ใช้จริงที่ไหน

middleware ที่ห่อ handler, functional options, memoization และ generator

:::caution[ข้อควรระวัง]
- ตัวแปรที่ถูกจับอาจถูกแชร์กับ goroutine อื่น ถ้ามีมากกว่าหนึ่ง goroutine เขียนต้องป้องกันด้วย mutex
- ก่อน Go 1.22 closure ใน loop จับ loop variable ตัวเดียวกันหมด
:::
