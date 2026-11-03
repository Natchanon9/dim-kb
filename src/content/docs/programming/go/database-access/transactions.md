---
title: "Transactions"
description: "BeginTx, Commit and deferred Rollback"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

A transaction makes several statements succeed or fail together. The standard shape is `BeginTx`, `defer Rollback`, then `Commit` at the end.

## Example

```go title="transfer.go"
func transfer(ctx context.Context, db *sql.DB, from, to int64, amount int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // no-op after a successful Commit

	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from); err != nil {
		return fmt.Errorf("debit: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to); err != nil {
		return fmt.Errorf("credit: %w", err)
	}
	return tx.Commit()
}
```

Pass `&sql.TxOptions{Isolation: sql.LevelSerializable}` when you need stricter isolation, and be ready to retry on serialization errors.

:::caution[Pitfalls]
- Inside a transaction use `tx.`, not `db.`; a `db` call runs on another connection outside the transaction.
- Keep transactions short; don't call external APIs while holding one.
:::
