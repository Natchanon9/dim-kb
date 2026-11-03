---
title: "Prepared Statements"
description: "parse query ครั้งเดียวแล้วรันหลายครั้ง"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

prepared statement ถูก database parse และวางแผนครั้งเดียว แล้วรันด้วย argument ต่างกันได้หลายครั้ง เหมาะกับ query เดิมใน loop ที่วนถี่

## ตัวอย่าง

```go title="insert.go"
stmt, err := db.PrepareContext(ctx, `INSERT INTO users (name) VALUES ($1)`)
if err != nil {
	return err
}
defer stmt.Close()
for _, n := range names {
	if _, err := stmt.ExecContext(ctx, n); err != nil {
		return err
	}
}
```

## จำเป็นไหม?

`db.QueryContext(ctx, query, args...)` ที่ใช้ placeholder ปลอดภัยจาก SQL injection อยู่แล้ว และ driver อย่าง pgx cache statement ให้อัตโนมัติ prepare เองเฉพาะใน loop ที่วนถี่ ส่วนการ insert ทีละมากๆ ใช้ `INSERT` หลายแถวหรือ `COPY` เร็วกว่าอีก

:::caution[ข้อควรระวัง]
- `*sql.Stmt` ที่ได้จาก `db` จะ prepare ใหม่ทุกครั้งที่ได้ connection ใหม่ ใน transaction ให้ใช้ `tx.StmtContext`
:::
