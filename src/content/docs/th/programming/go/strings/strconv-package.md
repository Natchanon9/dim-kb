---
title: "strconv Package"
description: "แปลงระหว่าง string กับตัวเลขหรือ boolean"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`strconv` แปลง string ไปเป็นตัวเลขหรือ boolean และกลับกัน ถ้า input ผิดจะคืน error

## Parse และ format

```go title="strconv.go"
n, _ := strconv.Atoi("42")
f, _ := strconv.ParseFloat("3.5", 64)
b, _ := strconv.ParseBool("true")
fmt.Println(n+1, f*2, b, strconv.Itoa(7)+"!", strconv.FormatFloat(1.5, 'f', 2, 64), strconv.Quote("hi"))
// 43 7 true 7! 1.50 "hi"

_, err := strconv.Atoi("12a")
// strconv.Atoi: parsing "12a": invalid syntax
```

| จาก string | เป็น string |
| --- | --- |
| `Atoi`, `ParseInt(s, base, bits)` | `Itoa`, `FormatInt(n, base)` |
| `ParseFloat(s, 64)` | `FormatFloat(f, 'f', prec, 64)` |
| `ParseBool` | `FormatBool` |
| `Unquote` | `Quote` |

:::caution[ข้อควรระวัง]
- เช็ค error ของ `Parse*`/`Atoi` เสมอ ถ้าไม่สน input ผิดจะกลายเป็น 0
- `ParseInt` ที่ `bitSize 32` ก็ยังคืน `int64` ต้องแปลงเองหลังเช็คแล้ว
:::
