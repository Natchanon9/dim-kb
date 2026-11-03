---
title: "Buffered Channels"
description: "channel ที่มีความจุ ส่งได้เลยจนกว่า buffer จะเต็ม"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

buffered channel มีคิว ส่งได้ทันทีถ้ายังมีที่ว่าง และรับได้ทันทีถ้ายังมีค่าค้างอยู่

## ตัวอย่าง

```go title="buffered.go"
ch := make(chan int, 2)
ch <- 1 // ยังมีที่ว่าง ไม่ต้องมีคนรับ
ch <- 2
fmt.Println(len(ch), cap(ch)) // 2 2
fmt.Println(<-ch, <-ch)       // 1 2 (FIFO)
```

## เมื่อไหร่ควรมี buffer

- **ขนาด 1** ให้ goroutine ส่งผลลัพธ์แล้วจบได้แม้ไม่มีใครอ่าน (กัน leak ตอน timeout)
- **รู้ขนาดงานแน่นอน** เช่น `make(chan Result, len(jobs))`
- **Semaphore**: `make(chan struct{}, n)` จำกัดงานพร้อมกันไว้ที่ n

อย่าเพิ่ม buffer ใหญ่ๆ เพื่อ "แก้" deadlock มักแค่ซ่อน bug ไว้จนกว่า load จะขึ้น

:::caution[ข้อควรระวัง]
- `len(ch)` เป็นแค่ภาพ ณ ขณะนั้น อาจเปลี่ยนก่อนเราใช้ อย่าเอาไปใช้ตัดสินใจ
:::
