---
title: "Named Return Values"
description: "ตั้งชื่อผลลัพธ์, naked return และแก้ผลลัพธ์ใน defer"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

ผลลัพธ์ตั้งชื่อได้ เริ่มต้นเป็น zero value และ `return` เปล่าๆ จะคืนค่าปัจจุบันของตัวแปรเหล่านั้น

## Naked return

```go title="split.go"
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // คืน x, y
}
// split(17) → 7 10
```

## ตรงไหนที่ตั้งชื่อแล้วได้ประโยชน์จริง

function ที่ defer ไว้อ่านและแก้ผลลัพธ์ที่มีชื่อได้ เป็นวิธีมาตรฐานในการเติม context ให้ทุก error ที่คืน:

```go title="parse.go"
func parse(s string) (n int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse %q: %w", s, err)
		}
	}()
	return strconv.Atoi(s)
}
// parse("x") → parse "x": strconv.Atoi: parsing "x": invalid syntax
```

ชื่อยังบอกด้วยว่าผลลัพธ์แต่ละตัวหมายถึงอะไร: `(width, height int)`

:::caution[ข้อควรระวัง]
- naked return ใน function ยาวๆ ทำให้ไม่รู้ว่าคืนอะไร ใช้กับ function สั้นๆ เท่านั้น
- `x, err := ...` ในบล็อกด้านในจะ shadow ผลลัพธ์ที่มีชื่อ ตัวที่ถูกคืนคือตัวนอก
:::
