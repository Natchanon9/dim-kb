package databaseaccess

import (
	"context"
	"database/sql"
)

func insertUsers(ctx context.Context, db *sql.DB, names []string) error {
	stmt, err := db.PrepareContext(ctx, `INSERT INTO users (name) VALUES ($1)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, n := range names {
		if _, err := stmt.ExecContext(ctx, n); err != nil { // parsed once, executed many times
			return err
		}
	}
	return nil
}

var _ = insertUsers
