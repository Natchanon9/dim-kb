---
title: "Struct Embedding"
description: "ประกอบ type ด้วย embedding โดย field และ method ถูก promote ขึ้นมา"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

embedding คือการใส่ type ไว้ใน struct โดยไม่ตั้งชื่อ field field และ method ของมันจะถูก **promote** ใช้ได้เหมือนเป็นของ struct ตัวนอก นี่คือ composition ไม่ใช่ inheritance

## ตัวอย่าง

```go title="embed.go"
type Base struct{ ID int }

func (b Base) Describe() string { return fmt.Sprintf("id=%d", b.ID) }

type Product struct {
	Base // embed
	Name string
}

p := Product{Base{7}, "pen"}
fmt.Println(p.ID, p.Describe(), p.Base.ID) // 7 id=7 7
```

## ไม่ใช่ inheritance

- `Product` ไม่ใช่ `Base` ส่งไปในที่ที่ต้องการ `Base` ไม่ได้
- method ที่ promote มายังได้ receiver เป็น `Base` ข้างใน ไม่ใช่ `Product`
- type ตัวนอก **shadow** field หรือ method ที่ promote มาได้ โดยประกาศของตัวเองชื่อเดียวกัน

การใช้จริงที่เจอบ่อยคือ embed interface ใน interface และ embed `sync.Mutex` ใน struct

:::caution[ข้อควรระวัง]
- embed type แล้ว method ทั้งชุดของมันจะเปิดออกจาก type ของเรา embed เฉพาะที่อยากให้เป็น public
- embed สอง type ที่มี field ชื่อเดียวกันทำให้ชื่อนั้นกำกวม ต้องเข้าถึงผ่านชื่อ type
:::
