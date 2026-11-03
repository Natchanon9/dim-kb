---
title: "If-Else"
description: "เงื่อนไขที่ไม่ต้องมีวงเล็บ และการ return ก่อน"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`if` รับเงื่อนไขแบบ boolean โดยไม่ต้องมีวงเล็บ แต่ต้องมีปีกกาเสมอ

## Syntax

```go title="grade.go"
func grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	}
	return "C"
}
// grade(85), grade(72), grade(10) → A B C
```

## Return ก่อนแทนการใช้ `else`

สไตล์ Go จะจัดการ error หรือกรณีพิเศษก่อนแล้ว return ทางหลักของโค้ดจะได้ไม่ต้องย่อหน้าลึก:

```go
if err != nil {
	return err
}
// ทางหลักทำต่อตรงนี้
```

:::caution[ข้อควรระวัง]
- เงื่อนไขต้องเป็น `bool`: `if n {}` ที่ n เป็น int compile ไม่ผ่าน
- `else` ต้องอยู่บรรทัดเดียวกับ `}` ที่ปิด
:::
