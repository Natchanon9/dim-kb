---
title: "sync Package"
description: "ภาพรวมเครื่องมือ synchronize ทั้งหมด"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`sync` รวมเครื่องมือระดับล่างสำหรับประสานงานระหว่าง goroutine ส่วนใหญ่มีหน้าของตัวเองในหมวด Goroutines และ Concurrency Patterns หน้านี้เป็นแผนที่รวม

## กล่องเครื่องมือ

| Type | ใช้ทำอะไร |
| --- | --- |
| `Mutex`, `RWMutex` | ป้องกัน state ที่แชร์ |
| `WaitGroup` | รอ goroutine |
| `Once`, `OnceFunc`, `OnceValue(s)` | รันครั้งเดียวแบบ lazy |
| `Pool` | ใช้ object ชั่วคราวซ้ำ |
| `Map` | map ที่ใช้พร้อมกันได้ สำหรับบางรูปแบบ |
| `Cond` | รอเงื่อนไข (ใช้น้อย channel มักอ่านง่ายกว่า) |

## sync.Map

```go title="syncmap.go"
var m sync.Map
m.Store("a", 1)
v, ok := m.Load("a")                    // 1 true
actual, loaded := m.LoadOrStore("a", 2) // 1 true (เก็บค่าเดิมไว้)
```

เร็วกว่า `map` + `Mutex` ก็ต่อเมื่อ key ถูกเขียนครั้งเดียวแล้วอ่านหลายครั้ง หรือแต่ละ goroutine แตะ key คนละชุดกัน นอกนั้นใช้ map ธรรมดากับ mutex: มี type และง่ายกว่า

:::caution[ข้อควรระวัง]
- ทุกตัวในนี้ห้าม copy หลังใช้งานครั้งแรก ให้ส่งเป็น pointer และรัน `go vet`
:::
