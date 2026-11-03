---
title: "new vs make"
description: "new จองค่าที่เป็นศูนย์ ส่วน make สร้าง slice, map และ channel"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

ทั้งคู่จอง memory แต่ใช้กับคนละอย่าง

| | `new(T)` | `make(T, args)` |
| --- | --- | --- |
| ใช้กับ | ทุก type | slice, map, chan เท่านั้น |
| คืน | `*T` ที่ชี้ไปค่าศูนย์ | `T` ที่พร้อมใช้ (ไม่ใช่ pointer) |
| ใช้บ่อยแค่ไหน | น้อย `&T{}` นิยมกว่า | ใช้กับ map และ channel เสมอ |

## ตัวอย่าง

```go title="newmake.go"
p := new(int) // *int ที่ชี้ไป int ค่าศูนย์
*p = 7
s := make([]int, 2)
m := make(map[string]int)
ch := make(chan int, 1)
m["a"] = 1
ch <- 9
fmt.Println(*p, s, m, <-ch) // 7 [0 0] map[a:1] 9

ps := new([]int) // pointer ไปยัง nil slice: แทบไม่เคยเป็นสิ่งที่ต้องการ
fmt.Println(*ps == nil) // true
```

:::caution[ข้อควรระวัง]
- `new(map[string]int)` ได้ pointer ไปยัง nil map เขียนแล้ว panic ให้ใช้ `make`
:::
