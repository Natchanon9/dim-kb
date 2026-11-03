---
title: "Blank Identifier"
description: "ตัวแทน _ สำหรับทิ้งค่า import และเช็คตอน compile"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

`_` เป็นตัวแทนที่เขียนได้อย่างเดียว กำหนดค่าให้มันคือการทิ้งค่านั้น compiler จะไม่บ่นเรื่องผลลัพธ์ที่ไม่ได้ใช้

## ใช้ทำอะไร

```go title="blank.go"
n, _ := strconv.Atoi("42") // ตั้งใจไม่สน error
for _, v := range items {}  // ไม่สน index

import _ "image/png" // import เพื่อให้ init ทำงานอย่างเดียว

var _ fmt.Stringer = celsius(0) // เช็คตอน compile ว่า celsius implement Stringer
```

บรรทัดสุดท้ายจะ compile ไม่ผ่านทันทีถ้า `celsius` ไม่ตรงกับ interface อีกต่อไป

:::caution[ข้อควรระวัง]
- การทิ้ง error ด้วย `_` ควรเป็นการตัดสินใจที่ตั้งใจและเกิดไม่บ่อย reviewer จะถามว่าทำไม
- อ่านค่าจาก `_` ไม่ได้: `x := _` compile ไม่ผ่าน
:::
