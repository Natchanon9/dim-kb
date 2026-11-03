---
title: "sqlx Package"
description: "scan แถวลง struct และขยาย IN clause"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`github.com/jmoiron/sqlx` ต่อยอด `database/sql` ให้ scan ลง struct ตาม tag `db`, ใช้ named parameter และขยาย `IN` ได้ โดยยังเขียน SQL เองเหมือนเดิม

## Scan ลง struct

```go title="users.go"
type userRow struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
}

db := sqlx.MustConnect("pgx", dsn)
var users []userRow
err := db.SelectContext(ctx, &users, `SELECT id, email FROM users WHERE active`)

var u userRow
err = db.GetContext(ctx, &u, `SELECT id, email FROM users WHERE id = $1`, id)
```

## IN clause

```go title="in.go"
q, args, _ := sqlx.In(`SELECT id FROM users WHERE id IN (?)`, []int{1, 2, 3})
q = sqlx.Rebind(sqlx.DOLLAR, q) // หรือ db.Rebind(q)
// SELECT id FROM users WHERE id IN ($1, $2, $3) [1 2 3]
```

## อยู่ตรงไหนเมื่อเทียบกับตัวอื่น

| เครื่องมือ | สไตล์ |
| --- | --- |
| `database/sql` | `Scan` เอง |
| `sqlx` | SQL + จับคู่กับ struct |
| `sqlc` | generate โค้ด Go จากไฟล์ SQL |
| GORM | ORM เต็มรูปแบบ |

:::caution[ข้อควรระวัง]
- `Select` โหลดทุกแถวเข้า memory ผลลัพธ์ใหญ่ให้ใช้ `Queryx` + `StructScan`
:::
