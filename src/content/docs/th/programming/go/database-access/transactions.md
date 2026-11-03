---
title: "Transactions"
description: "BeginTx, Commit และ Rollback ที่ defer ไว้"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

transaction ทำให้หลาย statement สำเร็จหรือล้มเหลวไปพร้อมกัน รูปแบบมาตรฐานคือ `BeginTx`, `defer Rollback` แล้ว `Commit` ตอนท้าย

## ตัวอย่าง

```go title="transfer.go"
func transfer(ctx context.Context, db *sql.DB, from, to int64, amount int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // ไม่มีผลถ้า Commit สำเร็จแล้ว

	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from); err != nil {
		return fmt.Errorf("debit: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to); err != nil {
		return fmt.Errorf("credit: %w", err)
	}
	return tx.Commit()
}
```

ส่ง `&sql.TxOptions{Isolation: sql.LevelSerializable}` เมื่อต้องการ isolation ที่เข้มขึ้น และเตรียม retry เมื่อเจอ serialization error

:::caution[ข้อควรระวัง]
- ใน transaction ให้ใช้ `tx.` ไม่ใช่ `db.` การเรียกผ่าน `db` จะไปรันบน connection อื่นนอก transaction
- ทำ transaction ให้สั้น อย่าเรียก API ภายนอกระหว่างที่ถือ transaction อยู่
:::
