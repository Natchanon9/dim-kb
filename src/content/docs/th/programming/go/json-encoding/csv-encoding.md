---
title: "CSV Encoding"
description: "อ่านและเขียน CSV ด้วย encoding/csv"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`encoding/csv` จัดการเครื่องหมายคำพูดและ comma ใน field ให้ ซึ่ง `strings.Split` ธรรมดาทำผิด

## ตัวอย่าง

```go title="csv.go"
r := csv.NewReader(strings.NewReader("name,age\nann,30\n\"b, c\",40\n"))
rows, _ := r.ReadAll()
fmt.Println(len(rows), rows[2][0]) // 3 b, c

var sb strings.Builder
w := csv.NewWriter(&sb)
w.Write([]string{"x", "y,z"})
w.Flush() // x,"y,z"
```

ไฟล์ใหญ่ให้เรียก `r.Read()` ใน loop แทน `ReadAll` ตั้ง `r.Comma = ';'` สำหรับไฟล์ที่คั่นด้วย semicolon และ `r.FieldsPerRecord = -1` ถ้าแต่ละแถวมี field ไม่เท่ากัน

:::caution[ข้อควรระวัง]
- `w.Flush()` และเช็ค `w.Error()` เสมอ
- CSV ที่ export จาก Excel อาจขึ้นต้นด้วย UTF-8 BOM ต้องตัดออกจาก header ตัวแรก
:::
