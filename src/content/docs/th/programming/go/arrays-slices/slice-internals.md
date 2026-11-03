---
title: "Slice Internals"
description: "slice header: pointer, length และ capacity"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

ค่าของ slice คือ header เล็กๆ ที่มี 3 field: **pointer** ชี้ไปที่ backing array, **length** และ **capacity** การ copy slice คือ copy header ไม่ใช่ข้อมูล

## แชร์ backing array

```go title="internals.go"
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3] // ชี้ไปที่ arr[1], len 2, cap 4
fmt.Println(s, len(s), cap(s)) // [2 3] 2 4
s[0] = 20 // เขียนลง arr
fmt.Println(arr) // [1 20 3 4 5]
```

capacity นับจากจุดเริ่มของ slice ไปจนสุด backing array `cap(s)` จึงเป็น 4

## การขยาย

ถ้า `append` เกิน capacity Go จะจอง array ที่ใหญ่กว่า (slice เล็กขยายประมาณ 2 เท่า slice ใหญ่ขยายน้อยกว่านั้น) copy element ไป แล้วคืน header ที่ชี้ไป array ใหม่

:::caution[ข้อควรระวัง]
- หลังจองใหม่ slice เก่ากับใหม่ไม่แชร์ข้อมูลกันแล้ว แต่ก่อนจองยังแชร์อยู่ โค้ดที่พึ่งพฤติกรรมแบบใดแบบหนึ่งจะเปราะ
- sub-slice เล็กๆ ของ slice ใหญ่ทำให้ backing array ทั้งก้อนยังอยู่ใน memory ถ้าจะเก็บนานให้ copy ออกด้วย `slices.Clone`
:::
