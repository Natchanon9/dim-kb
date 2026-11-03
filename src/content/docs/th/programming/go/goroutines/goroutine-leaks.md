---
title: "Goroutine Leaks"
description: "goroutine ที่ไม่เคยจบ และการให้ทางออกกับมัน"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

leak คือ goroutine ที่ถูกบล็อกตลอดไป: รอ channel ที่ไม่มีใครใช้แล้ว หรือวน loop โดยไม่มีทางออก แต่ละตัวถือ memory และทุกอย่างที่มันอ้างถึงไว้

## ให้ทุก goroutine มีทางออก

```go title="leak.go"
func counter(ctx context.Context, out chan<- int) {
	for i := 0; ; i++ {
		select {
		case out <- i:
		case <-ctx.Done(): // ถ้าไม่มี case นี้ goroutine จะบล็อกตลอดไป
			return
		}
	}
}

ctx, cancel := context.WithCancel(context.Background())
out := make(chan int)
go counter(ctx, out)
fmt.Println(<-out, <-out) // 0 1
cancel() // counter return
```

## สาเหตุที่พบบ่อย

- ส่งลง unbuffered channel หลังฝั่งรับเลิกรอไปแล้ว (เจอบ่อยกับ timeout) แก้ด้วย buffer 1 ช่อง หรือ `select` กับ `ctx.Done()`
- range บน channel ที่ไม่มีใครปิด
- ลืมเรียก `cancel()` หรือลืม stop `time.Ticker`

ตรวจหาได้ด้วย `runtime.NumGoroutine()` ใน test, goroutine profile ของ pprof หรือ `go.uber.org/goleak`

:::caution[ข้อควรระวัง]
- จำนวน goroutine ที่ค่อยๆ เพิ่มขึ้นใน metrics ของ production คืออาการคลาสสิก
:::
