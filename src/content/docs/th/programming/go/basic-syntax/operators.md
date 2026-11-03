---
title: "Operators"
description: "ตัวดำเนินการคณิตศาสตร์ เปรียบเทียบ ตรรกะ และ bitwise"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

operator ของ Go หน้าตาเหมือนภาษา C แต่มีเรื่องให้งงน้อยกว่า: ไม่มี ternary และ `++`/`--` เป็น statement

## คณิตศาสตร์และ bitwise

```go title="ops.go"
fmt.Println(7/2, 7%2, 7.0/2)           // 3 1 3.5
fmt.Println(1<<3, 6&3, 6|3, 6^3, 6&^3) // 8 2 7 5 4
x := 5
x += 2
x++ // เป็น statement ไม่ใช่ expression
fmt.Println(x) // 8
```

| กลุ่ม | Operator |
| --- | --- |
| คณิตศาสตร์ | `+ - * / %` (`/` ของ integer ตัดเศษ) |
| เปรียบเทียบ | `== != < <= > >=` |
| ตรรกะ | `&& \|\| !` (short-circuit) |
| Bitwise | `& \| ^ &^ << >>` (`&^` = AND NOT ใช้ล้าง bit) |
| กำหนดค่า | `= += -= *= /= %= <<= >>= &= \|= ^=` |

:::caution[ข้อควรระวัง]
- ไม่มี ternary `?:` ให้เขียน `if`
- `y := x++` compile ไม่ผ่าน
- หาร integer ด้วยตัวแปรที่เป็นศูนย์จะ panic ตอน runtime
:::
