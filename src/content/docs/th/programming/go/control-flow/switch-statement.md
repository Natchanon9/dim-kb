---
title: "Switch"
description: "switch ที่ไม่ต้อง break, หลายค่าต่อ case, switch แบบไม่มี tag และ fallthrough"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`switch` ของ Go หยุดที่ case แรกที่ตรง จึงไม่ต้องใส่ `break` และมักอ่านง่ายกว่า `if-else` ต่อกันยาวๆ

## Switch ตามค่า

```go title="switch.go"
func dayKind(day string) string {
	switch day {
	case "sat", "sun": // หลายค่าใน case เดียว
		return "weekend"
	default:
		return "weekday"
	}
}
```

## Switch แบบไม่มี tag และ fallthrough

```go title="tagless.go"
x := 15
switch { // เหมือน switch true: case แรกที่เป็นจริงชนะ
case x > 10:
	fmt.Println("big")
	fallthrough // รัน body ของ case ถัดไปโดยไม่เช็คเงื่อนไข
case x > 100:
	fmt.Println("huge?")
default:
	fmt.Println("small")
}
// big
// huge?
```

:::caution[ข้อควรระวัง]
- `fallthrough` ข้ามเงื่อนไขของ case ถัดไปไปเลย ซึ่งแทบไม่เคยเป็นสิ่งที่ต้องการ
- `break` ใน `switch` ที่อยู่ใน `for` ออกแค่ switch ถ้าจะออกจาก loop ต้องใช้ label
:::
