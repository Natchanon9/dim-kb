---
title: "context.WithCancel"
description: "context ที่สั่งหยุดได้ด้วยการเรียก cancel()"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`WithCancel` คืน context ลูกกับ function `cancel` เรียก `cancel` แล้ว `Done()` ของ context นั้นและลูกทั้งหมดจะถูกปิด

## ตัวอย่าง

```go title="cancel.go"
ctx, cancel := context.WithCancel(context.Background())
cancel()
<-ctx.Done()
fmt.Println(ctx.Err(), errors.Is(ctx.Err(), context.Canceled)) // context canceled true

ctx2, cancel2 := context.WithCancelCause(context.Background()) // Go 1.20+
cancel2(errors.New("user logged out"))
fmt.Println(context.Cause(ctx2)) // user logged out
```

ใช้หยุด worker ที่ทำงานเบื้องหลังตอน shutdown หรือหยุด goroutine พี่น้องเมื่อมีตัวหนึ่งได้คำตอบแล้ว

:::caution[ข้อควรระวัง]
- เรียก `cancel` เสมอ ส่วนใหญ่ใช้ `defer cancel()` แม้งานจะจบปกติ ไม่อย่างนั้น resource ของ context จะค้างจนกว่า context แม่จะจบ (`go vet` เตือน)
- เรียก `cancel` ซ้ำได้อย่างปลอดภัย
:::
