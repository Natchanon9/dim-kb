---
title: "context.WithTimeout"
description: "ยกเลิกเองอัตโนมัติเมื่อครบระยะเวลา"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`WithTimeout(parent, d)` ยกเลิกตัวเองหลังผ่านไป `d` เป็นวิธีมาตรฐานในการจำกัดเวลาของการเรียกออกไปข้างนอก

## ตัวอย่าง

```go title="timeout.go"
ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
defer cancel()
err := slowOp(ctx)
errors.Is(err, context.DeadlineExceeded) // true
```

## ที่ใช้บ่อย

```go
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()
row := db.QueryRowContext(ctx, "SELECT ...")
```

ใช้แบบเดียวกันกับ `http.NewRequestWithContext`, การเรียก gRPC และ Redis client

:::caution[ข้อควรระวัง]
- timeout จะได้ผลก็ต่อเมื่อโค้ดที่ถูกเรียกคอยดู `ctx.Done()` library ทำให้อยู่แล้ว แต่ loop ที่เขียนเองต้องทำเอง
- timeout ของลูกยาวเกินของแม่ไม่ได้ deadline ที่มาก่อนชนะ
:::
