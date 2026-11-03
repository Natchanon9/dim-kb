---
title: "Slice Gotchas"
description: "backing array ที่แชร์กัน append ที่ทำให้งง และสำเนาใน range"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

bug ของ slice ส่วนใหญ่มาจาก slice สองตัวแชร์ backing array เดียวกันโดยไม่รู้ตัว

## append สองครั้ง ช่องเดียวกัน

```go title="gotcha.go"
a := make([]int, 3, 4)
b := append(a, 1) // ยังพอใน cap: เขียนลง backing array ของ a
c := append(a, 2) // ช่องเดิมอีกรอบ
fmt.Println(b[3], c[3]) // 2 2  (b เปลี่ยนตามด้วย!)

safe := append(a[:3:3], 1) // full slice expression จำกัด cap บังคับให้ copy
fmt.Println(safe[3], b[3]) // 1 2
```

`s[low:high:max]` ตั้ง capacity เป็น `max-low` ทำให้ `append` ครั้งถัดไปต้องจองใหม่

## เช็คลิสต์

- **function คืน sub-slice** ของ input: ผู้เรียกแต่ละคนอาจทำข้อมูลของกันและกันเสีย ให้คืน `slices.Clone`
- **`range` ให้สำเนา:** แก้ที่ `s[i]` ไม่ใช่ค่าจาก loop
- **slice เล็กจาก array ใหญ่** ทำให้ array ใหญ่ค้างอยู่ใน memory
- **append ลง slice ที่รับมา** อาจเปลี่ยนหรือไม่เปลี่ยนข้อมูลของผู้เรียก ขึ้นกับ capacity

:::caution[ข้อควรระวัง]
- bug พวกนี้เกิดแค่บางครั้ง (ขึ้นกับ capacity) เลยทำให้เกิดซ้ำได้ยาก ถ้าไม่ชัดว่าใครเป็นเจ้าของข้อมูลให้ copy
:::
