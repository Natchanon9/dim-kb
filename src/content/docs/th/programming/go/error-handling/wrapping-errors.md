---
title: "Wrapping Errors"
description: "เติม context ด้วย %w และเก็บ error เดิมไว้ข้างใน"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

การ wrap คือเติม context ให้ error โดยเก็บตัวเดิมไว้ข้างใน ผู้เรียกจึงยังตรวจได้ว่าจริงๆ เกิดอะไรขึ้น

## %w vs %v

```go title="wrap.go"
base := errors.New("connection refused")
wrapped := fmt.Errorf("query users: %w", base) // %w เก็บ base ไว้ข้างใน
flat := fmt.Errorf("query users: %v", base)    // %v copy แค่ข้อความ
fmt.Println(wrapped)                            // query users: connection refused
errors.Unwrap(wrapped) == base // true
errors.Unwrap(flat) == nil     // true
```

error ที่ถูก wrap ต่อกันเป็นสาย และ `errors.Is` กับ `errors.As` จะไล่ไปตามสายนั้น

## เมื่อไหร่ไม่ควร wrap

ใช้ `%v` เมื่อไม่อยากให้ผู้เรียกผูกกับ error ภายใน (เช่น ซ่อนว่าเราใช้ database driver ตัวไหนไว้หลังขอบเขตของ package)

:::caution[ข้อควรระวัง]
- `fmt.Errorf` ใส่ `%w` ได้หลายตัว (Go 1.20+) ผลลัพธ์จะเป็น tree แทนที่จะเป็นสาย
:::
