---
title: "cmp Package"
description: "Compare, Less, Ordered และ Or สำหรับการเรียงลำดับ"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`cmp` (Go 1.21+) มี helper สำหรับเปรียบเทียบที่ใช้คู่กับ `slices.SortFunc`

## ตัวอย่าง

```go title="cmp.go"
cmp.Compare(1, 2)       // -1
cmp.Less("b", "a")      // false
cmp.Or("", "default")   // "default": ค่าแรกที่ไม่เป็นศูนย์ (1.22+)

slices.SortFunc(staff, func(a, b employee) int {
	return cmp.Or( // ตามอายุ แล้วตามชื่อ
		cmp.Compare(a.Age, b.Age),
		cmp.Compare(a.Name, b.Name),
	)
})
// [{cy 25} {al 30} {bob 30}]
```

| ตัว | ใช้ทำอะไร |
| --- | --- |
| `cmp.Ordered` | constraint ของ type ที่ใช้ `< <= > >=` ได้ |
| `cmp.Compare(a, b)` | -1 / 0 / +1 จัดการ NaN แบบสม่ำเสมอ |
| `cmp.Less(a, b)` | น้อยกว่าแบบ boolean |
| `cmp.Or(vals...)` | ค่าแรกที่ไม่เป็นศูนย์: sort หลาย key, ค่าเริ่มต้น |

:::caution[ข้อควรระวัง]
- `cmp.Or` ประเมินทุก argument ไม่ได้หยุดก่อนถ้า argument หลังๆ แพง
:::
