---
title: "Range over Channel"
description: "รับค่าจนกว่า channel จะถูกปิด"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`for v := range ch` รับค่าไปเรื่อยๆ จน channel ถูกปิด **และ** ไม่มีค่าค้าง เป็นวิธีปกติในการกินข้อมูลจาก stream

## ตัวอย่าง

```go title="range.go"
ch := make(chan string, 3)
ch <- "a"
ch <- "b"
close(ch)
for s := range ch {
	fmt.Println(s) // a, b
}
v, ok := <-ch // ปิดแล้วและว่าง
fmt.Printf("%q %v\n", v, ok) // "" false
```

## กติกาการปิด

- ปิดเพื่อบอกว่า "ไม่มีค่าอีกแล้ว" ไม่ใช่เพื่อคืน resource
- **ฝั่งส่ง** เป็นคนปิด และปิดครั้งเดียว
- ถ้ามีหลายผู้ส่ง ให้ปิดหลังทุกตัวเสร็จ (WaitGroup + goroutine ตัวเดียวที่เป็นคนปิด)

:::caution[ข้อควรระวัง]
- `range` บน channel ที่ไม่มีใครปิดจะไม่จบ: goroutine leak
:::
