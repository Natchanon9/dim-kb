---
title: "GOMAXPROCS"
description: "จำนวน thread ที่รันโค้ด Go ได้พร้อมกัน"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`GOMAXPROCS` คือจำนวน P: จำนวน goroutine ที่รันโค้ด Go **แบบขนาน** ได้พร้อมกัน ค่าเริ่มต้นเท่ากับจำนวน CPU ที่ใช้ได้

## อ่านและตั้งค่า

```go title="procs.go"
prev := runtime.GOMAXPROCS(2) // ตั้งค่า แล้วคืนค่าเดิม
defer runtime.GOMAXPROCS(prev)
runtime.GOMAXPROCS(0) // 0 = ถามค่าอย่างเดียว → 2
```

หรือตั้ง environment variable `GOMAXPROCS=4`

## Container

ตั้งแต่ Go 1.25 ค่าเริ่มต้นเคารพขีดจำกัด CPU ของ container (cgroups) และอัปเดตเมื่อขีดจำกัดเปลี่ยน ในเวอร์ชันเก่าบน Kubernetes ค่าเริ่มต้นคือจำนวน CPU ของ node ทำให้โดน throttle ทีมต่างๆ จึงใช้ `go.uber.org/automaxprocs`

:::caution[ข้อควรระวัง]
- ตั้งเกินจำนวน CPU ไม่ได้เพิ่มการทำงานขนาน มีแต่เพิ่มการแย่งกัน
- goroutine ที่บล็อกรอ I/O ไม่นับรวม server ที่เน้น I/O จึงไม่ต้องตั้งค่าสูงขึ้น
:::
