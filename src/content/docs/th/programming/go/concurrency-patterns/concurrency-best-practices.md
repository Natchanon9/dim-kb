---
title: "Concurrency Best Practices"
description: "หลักที่ทำให้โค้ด concurrent ของ Go ถูกต้อง"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

bug ด้าน concurrency ส่วนใหญ่มาจากความเป็นเจ้าของที่ไม่ชัด และ goroutine ที่ไม่มีทางออก

## เช็คลิสต์

- **รู้ว่าทุก goroutine จบอย่างไร** ก่อนเริ่มมัน
- **ส่ง `context.Context`** เป็น parameter แรกของทุกอย่างที่บล็อก และเคารพ `ctx.Done()`
- **ฝั่งส่งเป็นคนปิด** channel ฝั่งรับไม่ปิด
- **เลือก API แบบ synchronous ก่อน** ให้ผู้เรียกเพิ่ม concurrency เอง อย่าเริ่ม goroutine ใน function ของ library แล้ว return ออกไป
- **จำกัดจำนวนงานพร้อมกัน** ด้วย pool, semaphore หรือ `errgroup.SetLimit`
- **วาง lock ไว้ข้างข้อมูล** และทำ critical section ให้สั้น
- **รัน `-race`** ใน CI
- **วัดก่อนปรับ:** เพิ่ม goroutine ไม่ช่วยงานที่ใช้ CPU เกินจำนวน core

:::caution[ข้อควรระวัง]
- ใช้ `time.Sleep` เพื่อ "รอ goroutine" คือ bug ให้ใช้ WaitGroup หรือ channel
:::
