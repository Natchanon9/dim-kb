---
title: "พื้นฐาน Interface"
description: "ชุดของ method signature type ไหนมี method ครบก็ใช้ได้"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

interface ระบุรายการ method signature type ไหนก็ตามที่มี method เหล่านั้นจะใช้เป็น interface นั้นได้ function เดียวจึงทำงานกับหลาย type ได้

## ตัวอย่าง

```go title="speaker.go"
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "woof" }

type Robot struct{ ID int }

func (r Robot) Speak() string { return fmt.Sprintf("beep %d", r.ID) }

for _, s := range []Speaker{Dog{}, Robot{7}} {
	fmt.Println(s.Speak()) // woof, beep 7
}
```

## ค่าของ interface เก็บอะไร

ค่าของ interface เป็นคู่ (dynamic type, value) จะเป็น `nil` ก็ต่อเมื่อทั้งสองส่วนว่าง

:::caution[ข้อควรระวัง]
- เรียก method บน interface ที่เป็น nil จะ panic
- interface ที่มี method เดียวตั้งชื่อลงท้ายด้วย `-er`: `Reader`, `Stringer`, `Speaker`
:::
