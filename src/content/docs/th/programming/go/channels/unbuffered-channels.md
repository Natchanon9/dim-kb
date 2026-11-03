---
title: "Unbuffered Channels"
description: "การส่งกับการรับมาเจอกัน: ส่งต่อแบบ synchronous"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

unbuffered channel (`make(chan T)`) ไม่มีคิว การส่งจะบล็อกจนกว่าฝั่งรับจะรับค่าไป ทั้งสอง goroutine จึงมาเจอกันตรงจังหวะนั้น

## ตัวอย่าง

```go title="handoff.go"
ch := make(chan int)
done := make(chan struct{})
go func() {
	v := <-ch
	fmt.Println("got", v) // got 42
	close(done)
}()
ch <- 42 // บล็อกจน goroutine รับไป
<-done
```

## ทำไมถึงมีประโยชน์

การส่งต่อแบบนี้เป็นการรับประกัน: เมื่อการส่ง return ฝั่งรับได้ค่าไปแล้วแน่นอน unbuffered channel จึงเหมาะกับการส่งสัญญาณและการส่งต่อความเป็นเจ้าของข้อมูล

| | Unbuffered | Buffered |
| --- | --- | --- |
| การส่งบล็อกจนกว่า | มีคนรับไป | มีที่ว่าง |
| รับประกันว่าถึงมือแล้วตอนส่งเสร็จ | ใช่ | ไม่ |

:::caution[ข้อควรระวัง]
- ส่งและรับบน unbuffered channel เดียวกันใน goroutine เดียวกันจะ deadlock
:::
