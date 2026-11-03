---
title: "Type Switch"
description: "แยกทางตาม type จริงของค่าใน interface"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

type switch เช็คว่าค่าใน interface เป็น concrete type อะไร และให้ตัวแปรเป็น type นั้นในแต่ละ case

## Syntax

```go title="describe.go"
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int %d", x+1) // ตรงนี้ x เป็น int
	case string:
		return "string " + x
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("other %T", x)
	}
}
// describe(1) → int 2, describe(2.5) → other float64
```

ถ้า case มีหลาย type (`case int, int64:`) `x` จะยังเป็น type ของ interface

:::caution[ข้อควรระวัง]
- ใช้ได้กับค่าที่เป็น interface เท่านั้น (`any`, `error`, `io.Reader` …) ใช้กับ concrete type ไม่ได้
- ใช้ method บน interface ดีกว่า type switch ใหญ่ๆ เพิ่ม type ใหม่จะได้ไม่ต้องไล่แก้ทุก switch
:::
