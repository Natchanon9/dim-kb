---
title: "Pipelines"
description: "ต่อหลาย stage เข้าด้วยกันด้วย channel"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

pipeline คือ stage ที่ต่อกันเป็นทอด แต่ละ stage รับจาก channel ขาเข้า ทำงานหนึ่งอย่าง แล้วส่งไป channel ขาออกที่ตัวเองเป็นเจ้าของและเป็นคนปิด

## ตัวอย่าง

```go title="pipeline.go"
func source(nums ...int) <-chan int { /* ส่ง nums แล้วปิด */ }

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func sum(in <-chan int) int { /* บวกทุกค่าจาก in */ }

fmt.Println(sum(square(source(1, 2, 3)))) // 14
```

ทุก stage รันพร้อมกัน: ขณะที่ `sum` บวกค่าหนึ่ง `square` ก็ทำค่าถัดไป stage ที่ช้าแตกเป็นหลาย goroutine (fan-out) ได้

:::caution[ข้อควรระวัง]
- ถ้าฝั่งรับหยุดก่อน stage ต้นน้ำจะบล็อกตลอดไป ให้ส่ง `ctx` ไปและ `select` กับ `ctx.Done()` ทุกครั้งที่ส่ง
:::
