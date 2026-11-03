---
title: "พื้นฐาน Function"
description: "ประกาศ function, parameter และการส่งค่าแบบ pass-by-value"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

function ประกาศด้วย `func` ตามด้วย parameter และ type ของผลลัพธ์ argument ทุกตัวส่งแบบ **by value** คือ function ได้สำเนาไป

## ประกาศและเรียกใช้

```go title="add.go"
func add(a, b int) int { return a + b } // a, b ใช้ type เดียวกัน

func double(n int) { n *= 2 } // n เป็นสำเนา

x := 5
double(x)
fmt.Println(add(2, 3), x) // 5 5  (x ไม่เปลี่ยน)
```

## Pass by value

ถ้าอยากให้ function แก้ตัวแปรของผู้เรียก ให้ส่ง pointer (`*int`) ส่วน slice, map และ channel มี pointer อยู่ข้างใน function จึงแก้ **element** ได้โดยไม่ต้องส่ง pointer

:::caution[ข้อควรระวัง]
- ไม่มีค่า default ของ parameter และไม่มี overloading ให้ใช้ชื่อแยก (`NewServer`, `NewServerWithTLS`) หรือส่ง struct ของ option
- parameter ที่ไม่ได้ใช้ไม่เป็นไร แต่ตัวแปร local ที่ไม่ได้ใช้ compile ไม่ผ่าน
:::
