---
title: "WaitGroup"
description: "รอ goroutine กลุ่มหนึ่งให้เสร็จ"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`sync.WaitGroup` นับ goroutine ที่ยังทำงาน เรียก `Add` ก่อนเริ่ม `Done` เมื่อแต่ละตัวเสร็จ และ `Wait` จะบล็อกจนตัวนับเป็นศูนย์

## ตัวอย่าง

```go title="wg.go"
var (
	wg    sync.WaitGroup
	mu    sync.Mutex
	total int
)
for _, n := range []int{1, 2, 3} {
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		total += n
		mu.Unlock()
	}()
}
wg.Wait()
fmt.Println(total) // 6
```

Go 1.25 เพิ่ม `wg.Go(func() { ... })` ที่ทำ `Add(1)`, `go` และ `Done` ให้ในตัว

:::caution[ข้อควรระวัง]
- เรียก `Add` **ก่อน** `go` ไม่ใช่ข้างใน goroutine ไม่อย่างนั้น `Wait` อาจ return เร็วเกินไป
- ส่ง WaitGroup เป็น pointer การ copy ทำให้ตัวนับพัง (`go vet` เตือน)
- WaitGroup ไม่เก็บ error ถ้าต้องการให้ใช้ `errgroup`
:::
