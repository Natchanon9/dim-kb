---
title: "พื้นฐาน Channel"
description: "ส่งและรับค่าระหว่าง goroutine"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

channel คือท่อที่มี type เชื่อมระหว่าง goroutine ฝั่งหนึ่งส่งด้วย `ch <- v` อีกฝั่งรับด้วย `<-ch` และการส่งรับนี้ทำให้ทั้งสองฝั่ง synchronize กัน

## ตัวอย่าง

```go title="ping.go"
ch := make(chan string)
go func() { ch <- "ping" }() // ส่ง
msg := <-ch                  // รับ: รอจนมีการส่ง
fmt.Println(msg)             // ping
```

## Operation

| Operation | Syntax | หมายเหตุ |
| --- | --- | --- |
| สร้าง | `make(chan T)` / `make(chan T, n)` | zero value คือ `nil` |
| ส่ง | `ch <- v` | บล็อกจนมีคนรับหรือ buffer มีที่ว่าง |
| รับ | `v := <-ch` / `v, ok := <-ch` | `ok` เป็น false เมื่อปิดแล้วและว่าง |
| ปิด | `close(ch)` | ฝั่งส่งเท่านั้นที่ควรปิด |

"อย่าสื่อสารด้วยการแชร์ memory แต่ให้แชร์ memory ด้วยการสื่อสาร" ส่งความเป็นเจ้าของข้อมูลผ่าน channel แทนการ lock ตัวแปรที่แชร์กัน

:::caution[ข้อควรระวัง]
- ส่งลง channel ที่ปิดแล้วจะ panic และปิดซ้ำก็ panic
- ทุก goroutine บล็อกอยู่ที่ channel โดยไม่มีคู่ → `all goroutines are asleep - deadlock!`
:::
