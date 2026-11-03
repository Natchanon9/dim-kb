package databaseaccess

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRow struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
}

func activeUsers(ctx context.Context, db *sqlx.DB) ([]userRow, error) {
	var users []userRow
	err := db.SelectContext(ctx, &users, `SELECT id, email FROM users WHERE active`) // scans into the slice
	return users, err
}

var _ = activeUsers

func Example_sqlxIn() {
	q, args, _ := sqlx.In(`SELECT id FROM users WHERE id IN (?)`, []int{1, 2, 3}) // expand a slice
	q = sqlx.Rebind(sqlx.DOLLAR, q)                                               // ? → $1 for Postgres
	fmt.Println(q, args)
	// Output: SELECT id FROM users WHERE id IN ($1, $2, $3) [1 2 3]
}
