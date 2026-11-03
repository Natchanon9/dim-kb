---
title: "Channel Directions"
description: "type ของ channel แบบส่งอย่างเดียวและรับอย่างเดียว"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

parameter ของ function จำกัด channel ให้ทำได้ทางเดียวได้ แล้ว compiler จะปฏิเสธ operation ที่ผิดทาง

## ตัวอย่าง

```go title="directions.go"
func produce(out chan<- int) { // ส่งอย่างเดียว
	for i := range 3 {
		out <- i
	}
	close(out)
}

func consume(in <-chan int) int { // รับอย่างเดียว
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}

ch := make(chan int)
go produce(ch)
fmt.Println(consume(ch)) // 3
```

| Type | ทำได้ |
| --- | --- |
| `chan T` | ทั้งส่งและรับ |
| `chan<- T` | ส่ง (และปิด) เท่านั้น |
| `<-chan T` | รับเท่านั้น |

channel สองทางแปลงเป็นทางเดียวได้อัตโนมัติ

:::caution[ข้อควรระวัง]
- ปิด channel แบบรับอย่างเดียว compile ไม่ผ่าน ซึ่งตั้งใจไว้แบบนั้น: ฝั่งส่งเท่านั้นที่ปิด
:::
