---
title: "filepath Package"
description: "สร้างและแยก path ของไฟล์ให้ใช้ได้ทุก OS"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`path/filepath` จัดการ path ของไฟล์ด้วยตัวคั่นที่ถูกต้องของแต่ละ OS (`/` หรือ `\`) อย่าสร้าง path ด้วยการต่อ string

## Function ที่ใช้บ่อย

```go title="paths.go"
p := filepath.Join("data", "users", "..", "orders", "2024.csv") // จัดการ ".." ให้ด้วย
filepath.Base(p) // 2024.csv
filepath.Ext(p)  // .csv
filepath.Dir(p)  // data/orders
filepath.Match("*.csv", "report.csv") // true
```

| Function | ใช้ทำอะไร |
| --- | --- |
| `Join`, `Clean` | สร้าง path ที่ normalize แล้ว |
| `Base`, `Dir`, `Ext` | แยกส่วน |
| `Abs`, `Rel` | path แบบเต็ม / แบบสัมพัทธ์ |
| `Glob`, `Match` | wildcard |
| `WalkDir` | ไล่ทุกไฟล์ใน tree |

`path` (ไม่มี `filepath`) ใช้กับ path ที่คั่นด้วย slash เสมอ เช่น URL และ path ของ `fs.FS`

:::caution[ข้อควรระวัง]
- เอา input จากผู้ใช้มาต่อเป็น path อาจหลุดออกนอกโฟลเดอร์ (`../../etc/passwd`) ใช้ `filepath.IsLocal` (Go 1.20+) หรือ `os.Root` (Go 1.24+) เพื่อให้อยู่ในโฟลเดอร์ที่กำหนด
:::
