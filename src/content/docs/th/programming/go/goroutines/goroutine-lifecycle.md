---
title: "Goroutine Lifecycle"
description: "goroutine เริ่ม ถูกบล็อก และจบอย่างไร และทำไม main ต้องรอ"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

goroutine รันจนกว่า function ของมันจะ return ฆ่าจากข้างนอกไม่ได้ ต้องบอกให้หยุดแล้วมัน return เอง

## รอให้เสร็จ

```go title="lifecycle.go"
done := make(chan struct{})
go func() {
	defer close(done) // ส่งสัญญาณ: เสร็จแล้ว
	fmt.Println("working")
}()
<-done // ถ้าไม่มีบรรทัดนี้ main อาจ return ก่อน
fmt.Println("finished")
```

## สถานะ

| สถานะ | ความหมาย |
| --- | --- |
| Runnable | รอ CPU |
| Running | กำลังรันบน thread |
| Waiting | ถูกบล็อกที่ channel, lock, I/O, sleep หรือ syscall |
| Dead | function return แล้ว |

เมื่อ `main` return โปรแกรมจะจบ และ goroutine อื่นทั้งหมดหยุดกลางทาง deferred call ของพวกมันไม่ได้รัน

:::caution[ข้อควรระวัง]
- วางแผนการจบของทุก goroutine ก่อนเริ่ม: ใครบอกให้หยุด (context, channel ที่ปิด) และใครรอมัน
:::
