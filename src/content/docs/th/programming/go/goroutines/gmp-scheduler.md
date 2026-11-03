---
title: "GMP Scheduler"
description: "runtime จัด goroutine ลงบน OS thread อย่างไร"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

runtime ของ Go จัดคิว goroutine เอง โดยกระจาย goroutine จำนวนมากลงบน OS thread ไม่กี่ตัว โมเดลมี 3 ส่วน

| ตัวอักษร | คือ | จำนวน |
| --- | --- | --- |
| **G** | Goroutine | หลักพันถึงหลักล้าน |
| **M** | Machine: OS thread | เพิ่มเมื่อ thread ถูกบล็อกใน syscall |
| **P** | Processor: สิทธิ์ในการรันโค้ด Go พร้อมคิวงานของตัวเอง | `GOMAXPROCS` |

## ทำงานอย่างไร

- M ต้องถือ P ถึงจะรัน G ได้ และแต่ละ P มีคิวของ G ที่พร้อมรัน
- P ที่ว่างจะ **ขโมย** งานจากคิวของ P อื่น
- เมื่อ G ถูกบล็อกใน syscall M ของมันจะถูกแยกออก แล้ว P ย้ายไปอยู่กับ M ตัวอื่น G ตัวอื่นจึงรันต่อได้
- การรอ network ไปอยู่ที่ netpoller แทนการบล็อก thread
- ตั้งแต่ Go 1.14 G ที่รันนานจะถูก **preempt** แบบ asynchronous loop เดียวจึงแย่ง CPU ไว้คนเดียวไม่ได้

```go title="inspect.go"
fmt.Println("CPUs:", runtime.NumCPU())
fmt.Println("Ps (GOMAXPROCS):", runtime.GOMAXPROCS(0))
fmt.Println("goroutines:", runtime.NumGoroutine())
```

:::caution[ข้อควรระวัง]
- แทบไม่ต้องจูน scheduler เอง ถ้า performance ดูแปลกให้ profile ก่อน (`go tool trace`, pprof)
:::
