---
title: "Query ข้อมูล"
description: "QueryRow, Query กับ rows.Next และ Exec"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

method 3 ตัวนี้ครอบคลุมทั้งอ่านและเขียน ใช้เวอร์ชัน `...Context` เสมอ request จะได้ยกเลิกได้

## แถวเดียว

```go title="get.go"
var u User
err := db.QueryRowContext(ctx,
	`SELECT id, name, email FROM users WHERE id = $1`, id,
).Scan(&u.ID, &u.Name, &u.Email)
if errors.Is(err, sql.ErrNoRows) {
	return User{}, ErrUserNotFound
}
```

## หลายแถว

```go title="list.go"
rows, err := db.QueryContext(ctx, `SELECT id, name FROM users ORDER BY id LIMIT $1`, limit)
if err != nil {
	return nil, err
}
defer rows.Close() // คืน connection
for rows.Next() {
	var u User
	if err := rows.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	users = append(users, u)
}
return users, rows.Err()
```

## เขียน

```go
res, err := db.ExecContext(ctx, `UPDATE users SET name = $1 WHERE id = $2`, name, id)
n, _ := res.RowsAffected()
```

:::caution[ข้อควรระวัง]
- อย่าสร้าง SQL ด้วย `fmt.Sprintf` ร่วมกับ input จากผู้ใช้ ให้ส่งค่าเป็น argument เพื่อกัน SQL injection
- ลืม `rows.Close()` ทำให้ connection รั่วจน pool หมด
:::
