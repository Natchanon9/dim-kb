---
title: "Value vs Pointer Receiver"
description: "เมื่อไหร่ method ได้สำเนา และเมื่อไหร่แก้ค่าจริงได้"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

value receiver ได้สำเนา ส่วน pointer receiver ได้ address จึงแก้ตัวจริงได้และไม่ต้อง copy

## ตัวอย่าง

```go title="counter.go"
type Counter struct{ n int }

func (c Counter) Value() int  { return c.n }
func (c *Counter) Inc()       { c.n++ } // แก้ Counter ของผู้เรียก
func (c Counter) IncOnCopy()  { c.n++ } // แก้สำเนา: ไม่มีผล

var c Counter
c.Inc() // Go ใส่ &c ให้เอง
c.Inc()
c.IncOnCopy()
fmt.Println(c.Value()) // 2
```

## เลือกแบบไหน

| ใช้ pointer receiver เมื่อ | ใช้ value receiver เมื่อ |
| --- | --- |
| method แก้ค่า receiver | type เล็กและไม่ถูกแก้ (`time.Time`, `Point`) |
| struct ใหญ่ | type เป็น map, func หรือ chan (เป็น reference อยู่แล้ว) |
| struct มี `sync.Mutex` | |

ให้สม่ำเสมอ: ถ้ามี method ไหนต้องใช้ pointer receiver ให้ทุก method ใช้ pointer receiver

:::caution[ข้อควรระวัง]
- copy struct ที่มี `sync.Mutex` คือ copy lock ไปด้วย `go vet` จะแจ้งเตือน
:::
