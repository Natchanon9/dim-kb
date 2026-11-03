---
title: "Method Sets"
description: "ค่าหรือ pointer มี method อะไรบ้าง และมีผลกับ interface อย่างไร"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

method set เป็นตัวตัดสินว่า type นั้นตรงกับ interface ไหน

| Type | Method set |
| --- | --- |
| `T` | method ที่เป็น value receiver |
| `*T` | method ทั้ง value **และ** pointer receiver |

## ทำไมถึงสำคัญ

```go title="sets.go"
type Shape interface{ Area() float64 }

type Square struct{ S float64 }

func (s *Square) Area() float64 { return s.S * s.S } // pointer receiver

var sh Shape = &Square{2} // Square{2} จะ compile ไม่ผ่าน
fmt.Println(sh.Area())    // 4
```

เรียก `sq.Area()` บนตัวแปร `Square` ที่มี address ได้ยังใช้ได้ เพราะ Go เอา address ให้เอง แต่เก็บค่า `Square` ลง interface ไม่ได้ เพราะ interface เก็บสำเนาที่ไม่มี address

:::caution[ข้อควรระวัง]
- error "Square does not implement Shape (method Area has pointer receiver)" แปลว่าให้ส่ง `&Square{}`
:::
