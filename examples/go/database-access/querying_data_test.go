package databaseaccess

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func getUser(ctx context.Context, db *sql.DB, id int64) (User, error) {
	var u User
	err := db.QueryRowContext(ctx,
		`SELECT id, name, email, created_at FROM users WHERE id = $1`, id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	return u, err
}

func listUsers(ctx context.Context, db *sql.DB, limit int) ([]User, error) {
	rows, err := db.QueryContext(ctx, `SELECT id, name, email, created_at FROM users ORDER BY id LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // releases the connection
	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err() // errors that ended the loop early
}

func renameUser(ctx context.Context, db *sql.DB, id int64, name string) error {
	res, err := db.ExecContext(ctx, `UPDATE users SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("rename %d: %w", id, ErrUserNotFound)
	}
	return nil
}

var _, _, _ = getUser, listUsers, renameUser
