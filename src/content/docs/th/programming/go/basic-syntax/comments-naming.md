---
title: "Comment และการตั้งชื่อ"
description: "doc comment, ชื่อที่ export และสไตล์การตั้งชื่อของ Go"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

ตัวพิมพ์ใหญ่เล็กกำหนดว่าใครมองเห็นชื่อนั้น และ comment เหนือการประกาศจะกลายเป็น documentation สไตล์การตั้งชื่อเป็นส่วนหนึ่งของวัฒนธรรมภาษา

## Exported vs unexported

```go title="naming.go"
// MaxItems เป็น exported: ขึ้นต้นด้วยตัวพิมพ์ใหญ่
const MaxItems = 10

// discount เป็น unexported: เรียกได้เฉพาะโค้ดใน package นี้
func discount(price float64) float64 { return price * 0.9 }
```

## หลักการตั้งชื่อ

- **camelCase / PascalCase** ไม่ใช้ snake_case
- **ตัวย่อใช้ตัวพิมพ์เดียวกันทั้งคำ:** `userID`, `HTTPServer` ไม่ใช่ `userId`
- **scope สั้นใช้ชื่อสั้น:** `i`, `r`, `ctx` ส่วนของระดับ package ใช้ชื่อยาวกว่า
- **ไม่ซ้ำชื่อ package:** `http.Server` ไม่ใช่ `http.HTTPServer`
- **getter ไม่ต้องมี `Get`:** `u.Name()` ไม่ใช่ `u.GetName()`

## Doc comment

comment ที่อยู่เหนือชื่อที่ export ให้ขึ้นต้นด้วยชื่อนั้น แล้ว `go doc` กับ pkg.go.dev จะแสดงให้

:::caution[ข้อควรระวัง]
- ใช้ `//` เป็นหลัก ส่วน `/* */` ใช้กับ doc ของ package หรือปิดโค้ดชั่วคราว
- lint tool จะเตือนชื่อที่ export แต่ไม่มี doc comment
:::
