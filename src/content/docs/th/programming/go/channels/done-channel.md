---
title: "Done Channel"
description: "ปิด channel เพื่อบอกให้ goroutine หยุด"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

การปิด channel ปลุกฝั่งรับ **ทุกตัว** พร้อมกัน channel `done` ใช้คุณสมบัตินี้เป็นสัญญาณ "หยุด" แบบ broadcast

## ตัวอย่าง

```go title="done.go"
func numbers(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

done := make(chan struct{})
nums := numbers(done)
fmt.Println(<-nums, <-nums) // 0 1
close(done)                 // goroutine return
```

`chan struct{}` ไม่ได้ส่งข้อมูล มีแค่สัญญาณ

## ปัจจุบัน: ใช้ context

`context.Context` คือรูปแบบมาตรฐานของ pattern นี้ (`ctx.Done()` ก็คือ done channel) และมี deadline กับการส่งต่อข้าม API ด้วย โค้ดใหม่ให้ใช้ context

:::caution[ข้อควรระวัง]
- ส่งค่าลง `done` จะมี goroutine เห็นแค่ตัวเดียว ต้องปิดถึงจะไปถึงทุกตัว
:::
