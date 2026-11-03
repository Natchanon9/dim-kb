---
title: "การสร้าง Error"
description: "errors.New สำหรับข้อความคงที่ fmt.Errorf สำหรับข้อความที่มีค่าแทรก"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

function สองตัวนี้ครอบคลุมการสร้าง error เกือบทั้งหมด

## ตัวอย่าง

```go title="create.go"
e1 := errors.New("not found")            // ข้อความคงที่
e2 := fmt.Errorf("user %d not found", 7) // ข้อความที่แทรกค่า
fmt.Println(e1) // not found
fmt.Println(e2) // user 7 not found
```

| ต้องการ | ใช้ |
| --- | --- |
| ข้อความคงที่ หรือทำ sentinel | `errors.New` |
| แทรกค่า | `fmt.Errorf("... %d", v)` |
| เติม context ให้ error อื่น | `fmt.Errorf("...: %w", err)` |
| มี field เสริมให้ผู้เรียกตรวจ | custom error type |

## สไตล์ข้อความ

ตัวพิมพ์เล็ก ไม่มีเครื่องหมายวรรคตอนท้าย ไม่ขึ้นต้นด้วย "error:" เพราะข้อความจะถูกต่อกันเป็นสาย: `load config: open file: permission denied`

:::caution[ข้อควรระวัง]
- เรียก `errors.New("x")` สองครั้งได้ error **คนละตัว** เทียบด้วย `==` จะได้ false
:::
