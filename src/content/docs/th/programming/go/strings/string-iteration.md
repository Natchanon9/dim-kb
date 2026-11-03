---
title: "วนลูป String"
description: "range ตาม rune vs index ตาม byte"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`for i, r := range s` ถอด UTF-8 แล้วให้ rune ส่วน `for i := 0; i < len(s); i++` แบบคลาสสิกให้ byte ดิบ

## ทั้งสองแบบ

```go title="iterate.go"
s := "aก"
for i, r := range s { // ทีละ rune; i คือตำแหน่ง byte
	fmt.Println(i, string(r), r) // 0 a 97, 1 ก 3585
}
for i := 0; i < len(s); i++ { // ทีละ byte
	fmt.Print(s[i], " ") // 97 224 184 129
}
```

กับข้อความใช้ `range` ส่วน loop ทีละ byte ใช้กับ protocol ที่เป็น ASCII หรือการ parse ข้อมูล binary เท่านั้น

:::caution[ข้อควรระวัง]
- index จาก `range` กระโดด (0, 1, 4 …) เพราะเป็นตำแหน่ง byte ไม่ใช่ตำแหน่งตัวอักษร
:::
