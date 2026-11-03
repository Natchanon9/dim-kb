---
title: "slices Package (Modern)"
description: "helper ที่ใช้ iterator ซึ่งเพิ่มใน Go 1.23"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go 1.23 เชื่อม `slices` เข้ากับ iterator ผสมกันแล้วแทน loop เล็กๆ ได้

## ตัวอย่าง

```go title="slices.go"
s := []int{3, 1, 2}
slices.Sorted(slices.Values(s))   // [1 2 3]
slices.Collect(Countdown(3))      // [3 2 1]
for chunk := range slices.Chunk([]int{1, 2, 3, 4, 5}, 2) {
	fmt.Println(chunk) // [1 2] [3 4] [5]
}
```

| Function | ทำอะไร |
| --- | --- |
| `All(s)` | iterator ของ index, value |
| `Values(s)` | iterator ของ value |
| `Backward(s)` | index, value จากท้าย |
| `Collect(seq)` | iterator → slice ใหม่ |
| `AppendSeq(s, seq)` | ต่อค่าจาก iterator |
| `Sorted(seq)`, `SortedFunc` | เก็บแล้ว sort |
| `Chunk(s, n)` | sub-slice ขนาด n |

function รุ่นเก่า (`Sort`, `Contains`, `Index` …) อยู่ที่ [slices Package](/dim-kb/th/programming/go/arrays-slices/slices-package/)

:::caution[ข้อควรระวัง]
- `Chunk` คืนมุมมองของ slice เดิม ไม่ใช่สำเนา
:::
