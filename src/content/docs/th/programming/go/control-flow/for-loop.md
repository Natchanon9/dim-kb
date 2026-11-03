---
title: "For Loop"
description: "for เป็น loop เดียวของ Go: แบบคลาสสิก, แบบ while, วนไม่รู้จบ และ range over int"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go มี keyword สำหรับ loop ตัวเดียวคือ `for` ครอบคลุมทุกแบบที่ภาษาอื่นแยกเป็น `for`, `while` และ `do`

## 4 รูปแบบ

```go title="loops.go"
sum := 0
for i := 0; i < 5; i++ { // แบบคลาสสิก 3 ส่วน
	sum += i
}
n := 1
for n < 100 { // แบบ while
	n *= 2
}
count := 0
for { // วนไม่รู้จบ
	count++
	if count == 3 {
		break
	}
}
for i := range 3 { // Go 1.22+: range over int → 0 1 2
	fmt.Print(i, " ")
}
fmt.Println(sum, n, count) // 10 128 3
```

## Loop variable (Go 1.22)

ตั้งแต่ Go 1.22 แต่ละรอบได้ loop variable ตัว **ใหม่** closure และ goroutine ที่สร้างใน loop จึงเห็นค่าของรอบตัวเอง

:::caution[ข้อควรระวัง]
- ถ้า Go < 1.22 (ดูบรรทัด `go` ใน `go.mod`) closure ใน loop จะใช้ตัวแปรตัวเดียวกัน โค้ดเก่าแก้ด้วย `i := i`
- ส่วนหัวของ loop ไม่มีวงเล็บ แต่ต้องมีปีกกา
:::
