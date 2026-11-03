---
title: "Querying Data"
description: "QueryRow, Query with rows.Next, and Exec"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Three methods cover reads and writes. Always use the `...Context` versions so requests can be cancelled.

## One row

```go title="get.go"
var u User
err := db.QueryRowContext(ctx,
	`SELECT id, name, email FROM users WHERE id = $1`, id,
).Scan(&u.ID, &u.Name, &u.Email)
if errors.Is(err, sql.ErrNoRows) {
	return User{}, ErrUserNotFound
}
```

## Many rows

```go title="list.go"
rows, err := db.QueryContext(ctx, `SELECT id, name FROM users ORDER BY id LIMIT $1`, limit)
if err != nil {
	return nil, err
}
defer rows.Close() // releases the connection
for rows.Next() {
	var u User
	if err := rows.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	users = append(users, u)
}
return users, rows.Err()
```

## Writes

```go
res, err := db.ExecContext(ctx, `UPDATE users SET name = $1 WHERE id = $2`, name, id)
n, _ := res.RowsAffected()
```

:::caution[Pitfalls]
- Never build SQL with `fmt.Sprintf` and user input; pass values as arguments to prevent SQL injection.
- Forgetting `rows.Close()` leaks connections until the pool runs dry.
:::
