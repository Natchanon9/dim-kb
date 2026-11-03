---
title: "การแปลง Type"
description: "แปลง type ด้วย T(x) และใช้ strconv กับ string"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Go ไม่แปลง type ให้อัตโนมัติ ต้องเขียน `T(x)` เองเวลาแปลงระหว่างตัวเลข และใช้ `strconv` กับ string

## ตัวเลข

```go title="convert.go"
i := 42
f := float64(i)
x := 3.9
fmt.Println(f, int(x)) // 42 3  (int() ตัดทศนิยมทิ้ง)
```

## String กับตัวเลข

```go title="strconv.go"
s := strconv.Itoa(i)          // "42"
n, err := strconv.Atoi("123") // 123, nil
fmt.Println(s, n, err, string(rune(65))) // 42 123 <nil> A
```

| ต้องการ | ใช้ |
| --- | --- |
| int ↔ string | `strconv.Itoa`, `strconv.Atoi` |
| float/bool ↔ string | `strconv.ParseFloat`, `strconv.FormatFloat`, `strconv.ParseBool` |
| ค่าอะไรก็ได้ → string | `fmt.Sprint(v)` (ช้ากว่า ใช้กับ log ได้) |

:::caution[ข้อควรระวัง]
- `string(65)` ได้ `"A"` ไม่ใช่ `"65"` และ `go vet` จะเตือน ให้ใช้ `strconv.Itoa`
- แปลงเป็น integer ที่เล็กกว่าจะวนค่าเงียบๆ: `uint8(300)` ของ constant compile ไม่ผ่าน แต่ของตัวแปรจะวนค่า
:::
