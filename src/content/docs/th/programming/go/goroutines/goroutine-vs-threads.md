---
title: "Goroutine vs Thread"
description: "ทำไม goroutine ถูกกว่า OS thread"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

goroutine ถูกจัดคิวโดย runtime ของ Go ไม่ใช่ OS จึงถูกกว่ามาก การสร้าง goroutine ต่อ request หรือต่องานจึงเป็นเรื่องปกติ

## เปรียบเทียบ

| | Goroutine | OS thread |
| --- | --- | --- |
| stack เริ่มต้น | ~2 KB ขยายและหดได้ | ~1–8 MB คงที่ |
| การสร้าง | ระดับ microsecond ไม่มี syscall | ต้อง syscall ช้ากว่ามาก |
| การสลับ | ทำใน user space ถูก | context switch ของ kernel |
| ตัวตน | ไม่เปิดเผย ID | มี thread ID |
| จำนวนทั่วไป | 100k+ | หลักพัน |

```go title="many.go"
var (
	wg sync.WaitGroup
	n  atomic.Int64
)
for range 10_000 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.Add(1)
	}()
}
wg.Wait()
fmt.Println(n.Load()) // 10000
```

:::caution[ข้อควรระวัง]
- ถูกไม่ได้แปลว่าฟรี: goroutine ไม่จำกัดจำนวนที่ยิง database หรือ API จะทำให้ปลายทางรับไม่ไหว ให้จำกัดด้วย worker pool หรือ semaphore
:::
