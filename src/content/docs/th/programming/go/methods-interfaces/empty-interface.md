---
title: "Empty Interface"
description: "any / interface{} เก็บค่าได้ทุก type"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`any` (alias ของ `interface{}` ตั้งแต่ Go 1.18) ไม่มี method ทุก type จึงตรงกับมัน เก็บค่าอะไรก็ได้

## ตัวอย่าง

```go title="any.go"
var v any = 42
v = "hi"
v = []int{1}
fmt.Printf("%T\n", v) // []int

values := []any{1, "a", true, nil}
fmt.Println(values...) // 1 a true <nil>
```

## ใช้เมื่อไหร่

- ข้อมูลที่ไม่รู้ type จริงๆ: decode JSON ที่ไม่รู้รูปแบบ (`map[string]any`), `fmt.Println`, field ของ log
- จะเอาค่ากลับมาใช้ type assertion หรือ type switch

ถ้ารู้ชุดของ type อยู่แล้ว ใช้ generics หรือ interface จริงดีกว่า

:::caution[ข้อควรระวัง]
- `any` ทิ้งการเช็ค type ตอน compile ความผิดพลาดจะไปโผล่เป็น panic ตอน runtime
- ตัวเลขใน JSON ที่ decode ลง `any` จะเป็น `float64` ไม่ใช่ `int`
:::
