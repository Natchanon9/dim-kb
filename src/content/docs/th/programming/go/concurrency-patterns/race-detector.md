---
title: "Race Detector"
description: "หา data race ด้วย -race"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

data race คือสอง goroutine แตะ memory เดียวกันในเวลาเดียวกัน โดยมีอย่างน้อยหนึ่งตัวเขียนและไม่มีการ synchronize ผลลัพธ์จะคาดเดาไม่ได้ race detector หาให้ตอน runtime

## race และวิธีแก้

```go title="race.go"
count := 0
for range 2 {
	go func() { count++ }() // RACE: อ่าน-แก้-เขียนโดยไม่ synchronize
}

var safe atomic.Int64
for range 2 {
	go func() { safe.Add(1) }() // แก้แล้ว
}
```

```sh
go test -race ./...
go run -race .
```

รายงานจะแสดง stack trace ทั้งสองฝั่ง: ตรงที่เขียน และตรงที่เข้าถึงชนกัน

## ใช้งาน

- ใช้กับ test ใน CI เสมอ
- ใช้ใน build ของ staging ถ้ารับ overhead CPU และ memory ประมาณ 5–10 เท่าได้

:::caution[ข้อควรระวัง]
- เจอเฉพาะ race ในโค้ดที่ถูกรันจริง ไม่มีรายงานไม่ได้พิสูจน์ว่าไม่มี race
:::
