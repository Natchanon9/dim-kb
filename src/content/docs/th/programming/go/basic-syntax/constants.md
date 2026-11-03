---
title: "Constants"
description: "ค่า const, constant แบบมี type และไม่มี type"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`const` ถูกกำหนดตายตัวตอน compile ทำเป็น constant ได้เฉพาะตัวเลข string และ boolean (รวม rune)

## ประกาศ constant

```go title="const.go"
const Pi = 3.14159 // untyped constant

const timeout = 5 * time.Second
var f32 float32 = Pi // Pi แบบ untyped ใส่ float type ไหนก็ได้
fmt.Println(Pi, f32, timeout) // 3.14159 3.14159 5s
```

## Untyped vs typed

- **Untyped** (`const Pi = 3.14`): มี type เริ่มต้น แต่ปรับตามที่นำไปใช้ และแม่นยำสูง
- **Typed** (`const Pi float64 = 3.14`): ทำตัวเหมือนค่า `float64` ทุกประการ

ใช้แบบ untyped เป็นหลัก ยกเว้นตอนทำ type แบบ enum

:::caution[ข้อควรระวัง]
- slice, map และ struct เป็น `const` ไม่ได้ ให้ใช้ `var` ระดับ package
- constant ที่ล้น type ปลายทาง compile ไม่ผ่าน ไม่ได้วนค่าตอน runtime
:::
