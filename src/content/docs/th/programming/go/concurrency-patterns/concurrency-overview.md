---
title: "ภาพรวม Concurrency"
description: "concurrency vs parallelism และควรหยิบเครื่องมือไหน"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

concurrency คือการจัดโครงโปรแกรมเป็นงานย่อยที่เป็นอิสระต่อกัน ส่วน parallelism คือการรันงานในเวลาเดียวกันจริงบนหลาย CPU Go ให้ concurrency มา และ runtime จะทำให้เกิด parallelism เองเมื่อมี CPU ว่าง

## เลือกเครื่องมือ

| ปัญหา | เครื่องมือ |
| --- | --- |
| รัน N งานแล้วรอ | `sync.WaitGroup` |
| รัน N งาน หยุดเมื่อเจอ error แรก | `errgroup` |
| ส่งข้อมูลหรือความเป็นเจ้าของระหว่าง goroutine | Channel |
| ป้องกัน state ที่แชร์ | `sync.Mutex` / `RWMutex` |
| ตัวนับหรือ flag ตัวเดียว | `sync/atomic` |
| initialize ครั้งเดียว | `sync.Once`, `sync.OnceValue` |
| ยกเลิกและตั้ง deadline | `context` |
| จำกัดจำนวนงานพร้อมกัน | worker pool หรือ semaphore channel |
| ใช้ object ชั่วคราวซ้ำ | `sync.Pool` |

## Channel หรือ mutex?

ใช้ channel เมื่อ goroutine ส่งงานต่อให้กัน (pipeline, pool, event) ใช้ mutex เมื่อหลาย goroutine แตะ state ก้อนเดียวกัน (cache, map ของตัวนับ) ทั้งคู่เป็นสไตล์ Go ปกติ เลือกตัวที่ทำให้โค้ดง่ายกว่า

:::caution[ข้อควรระวัง]
- concurrency เพิ่ม bug ทำเวอร์ชันแบบทีละขั้นให้ทำงานถูกก่อน แล้วค่อยเพิ่ม goroutine ตรงที่ profile แล้วเห็นว่ามีการรอ
:::
