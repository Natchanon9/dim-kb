---
title: "maps Package (Modern)"
description: "Keys, Values, All และ Collect (Go 1.23)"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Go 1.23 เพิ่ม function แบบ iterator ให้ `maps` บรรทัดเดียวที่ใช้บ่อยที่สุดคือ key ที่เรียงแล้ว

## ตัวอย่าง

```go title="maps.go"
m := map[string]int{"b": 2, "a": 1}
slices.Sorted(maps.Keys(m))          // [a b]
copyOfM := maps.Collect(maps.All(m)) // iterator → map ใหม่
```

| Function | ทำอะไร |
| --- | --- |
| `Keys(m)` | iterator ของ key |
| `Values(m)` | iterator ของ value |
| `All(m)` | iterator ของ key, value |
| `Collect(seq2)` | iterator แบบ key-value → map ใหม่ |
| `Insert(m, seq2)` | ใส่คู่ค่าลง map ที่มีอยู่ |

function ของ 1.21 (`Clone`, `Equal`, `DeleteFunc`, `Copy`) อยู่ที่ [maps Package](/dim-kb/th/programming/go/maps/maps-package/)

:::caution[ข้อควรระวัง]
- `maps.Keys` ใน standard library คืน iterator ไม่ใช่ slice ส่วนโค้ดเก่าที่ใช้ `golang.org/x/exp/maps.Keys` ได้ slice
:::
