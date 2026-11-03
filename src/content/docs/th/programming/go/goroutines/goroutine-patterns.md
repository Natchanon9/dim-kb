---
title: "Goroutine Patterns"
description: "generator และการจำกัดจำนวนงานพร้อมกัน"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

รูปแบบไม่กี่แบบครอบคลุมการใช้ goroutine เกือบทั้งหมด หลักที่อยู่เบื้องหลังคือ goroutine ที่สร้าง channel เป็นเจ้าของและเป็นคนปิดมัน

## Generator

```go title="gen.go"
func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range n {
			ch <- i
		}
	}()
	return ch
}

for v := range gen(3) { ... } // 0, 1, 2
```

## จำกัดจำนวนงานพร้อมกัน (semaphore)

```go title="sem.go"
sem := make(chan struct{}, 2) // พร้อมกันได้สูงสุด 2
for range 5 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{}        // จอง
		defer func() { <-sem }() // คืน
		doWork()
	}()
}
wg.Wait()
```

รูปแบบอื่น (worker pool, pipeline, fan-in/fan-out, errgroup) อยู่ใน [Channels](/dim-kb/th/programming/go/channels/channel-basics/) และ [Concurrency Patterns](/dim-kb/th/programming/go/concurrency-patterns/concurrency-overview/)

:::caution[ข้อควรระวัง]
- คืน channel แบบรับอย่างเดียว (`<-chan T`) ผู้เรียกจะได้ปิดหรือส่งลงไปไม่ได้
:::
