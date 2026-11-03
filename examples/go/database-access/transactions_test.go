package databaseaccess

import (
	"context"
	"database/sql"
	"fmt"
)

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

var _ = transfer
