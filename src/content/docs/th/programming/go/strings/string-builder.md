---
title: "String Builder"
description: "สร้าง string อย่างมีประสิทธิภาพด้วย strings.Builder"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

ต่อ string ด้วย `+` ใน loop จะ copy string ทั้งก้อนทุกรอบ ส่วน `strings.Builder` ต่อลง buffer ที่ขยายได้

## ตัวอย่าง

```go title="builder.go"
var sb strings.Builder // zero value ใช้ได้เลย
for i := range 3 {
	fmt.Fprintf(&sb, "%d,", i) // Builder เป็น io.Writer
}
sb.WriteString("done")
fmt.Println(sb.String(), sb.Len()) // 0,1,2,done 10
```

method: `WriteString`, `WriteByte`, `WriteRune`, `Write`, `Grow(n)` สำหรับจองขนาดก่อน และ `Reset`

## เมื่อไหร่ควรใช้

- ใช้ `+` ไม่กี่ครั้งนอก loop: ไม่เป็นไร
- ต่อ slice: `strings.Join`
- สร้างใน loop: `strings.Builder`

:::caution[ข้อควรระวัง]
- อย่า copy `Builder` หลังเขียนไปแล้ว ให้ส่ง `*strings.Builder`
:::
