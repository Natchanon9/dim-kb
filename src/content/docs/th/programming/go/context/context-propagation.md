---
title: "การส่งต่อ Context"
description: "ส่ง ctx ตัวเดียวกันผ่านทุกชั้นของ request"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

context ของ request ไหลจาก handler ลงไปผ่าน service และ repository จนถึง database หรือ HTTP client เมื่อ client ตัดการเชื่อมต่อหรือเลย deadline ทุกชั้นจะหยุด

## ตัวอย่าง

```go title="layers.go"
func repoFind(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("repo: %w", err)
	}
	return db.QueryRowContext(ctx, "...").Err()
}

func serviceGet(ctx context.Context) error {
	if err := repoFind(ctx); err != nil {
		return fmt.Errorf("service: %w", err)
	}
	return nil
}

// request ที่ถูกยกเลิก → service: repo: context canceled
```

## กติกา

- ทุก function ที่ทำ I/O รับ `ctx` เป็นตัวแรก แล้วส่งต่อไปตรงๆ หรือสร้างลูก (`WithTimeout`)
- ใช้เวอร์ชัน `...Context`: `QueryContext`, `NewRequestWithContext`
- งานที่ต้องอยู่ต่อหลัง request จบ (audit log, ส่ง email แบบ async) ใช้ `context.WithoutCancel(ctx)` (Go 1.21+) ค่าต่างๆ ยังอยู่แต่ไม่ถูกยกเลิกตาม

:::caution[ข้อควรระวัง]
- เริ่ม goroutine ด้วย `ctx` ของ request แล้ว return ไปเลย: งานจะถูกยกเลิกทันทีที่ handler return
:::
