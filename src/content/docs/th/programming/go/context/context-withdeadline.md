---
title: "context.WithDeadline"
description: "ยกเลิก ณ เวลาที่กำหนด"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`WithDeadline(parent, t)` ยกเลิกตอนเวลา `t` ส่วน `WithTimeout(parent, d)` ก็คือ `WithDeadline(parent, time.Now().Add(d))`

## ตัวอย่าง

```go title="deadline.go"
deadline := time.Now().Add(50 * time.Millisecond)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
d, ok := ctx.Deadline() // deadline, true

child, cancelChild := context.WithTimeout(ctx, time.Hour)
defer cancelChild()
cd, _ := child.Deadline() // ยังเป็น deadline ของแม่
```

## เมื่อไหร่เวลาตายตัวเหมาะกว่า

เมื่องานทั้งก้อนต้องเสร็จภายในเวลาเดียว (ปิดรอบ batch, SLA ของ request ที่มีหลายขั้นตอน) ให้คำนวณ deadline ครั้งเดียว แล้วส่ง context ตัวเดียวกันไปทุกขั้น

:::caution[ข้อควรระวัง]
- เช็คเวลาที่เหลือด้วย `time.Until(d)` ก่อนเริ่มงานที่ทำไม่ทันแน่ๆ
:::
