---
title: "Arrays"
description: "array ความยาวคงที่ และทำไมแทบไม่ได้ใช้ตรงๆ"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

array มีความยาวคงที่และความยาวเป็นส่วนหนึ่งของ type: `[3]int` กับ `[4]int` เป็นคนละ type array เป็นค่า การกำหนดค่าจึง copy ทุก element

## ประกาศและ copy

```go title="arrays.go"
var a [3]int
b := [...]string{"x", "y"} // ความยาวนับจาก literal
a[0] = 1
c := a // copy ทุก element
c[0] = 99
fmt.Println(a, len(b), c, a == [3]int{1, 0, 0})
// [1 0 0] 2 [99 0 0] true
```

## ใช้ array เมื่อไหร่

ส่วนใหญ่ใช้ slice ส่วน array เหมาะกับข้อมูลขนาดตายตัว เช่น ค่า SHA-256 (`[32]byte`), สี RGB หรือใช้เป็น key ของ map ที่เทียบค่าได้ (`map[[2]int]bool`)

:::caution[ข้อควรระวัง]
- ส่ง array ใหญ่เข้า function จะ copy ทั้งก้อน ให้ส่ง slice หรือ pointer แทน
- index เกินขอบเป็น compile error ถ้าเป็น constant และ panic ตอน runtime ถ้าไม่ใช่
:::
