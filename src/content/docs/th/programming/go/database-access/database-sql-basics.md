---
title: "พื้นฐาน database/sql"
description: "API เดียวสำหรับ SQL database และ driver ตามแต่ละ database"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`database/sql` เป็น interface กลางสำหรับ SQL database เพิ่ม **driver** ของ database ที่ใช้ แล้วโค้ดเดียวกันใช้ได้กับ PostgreSQL, MySQL หรือ SQLite

## Setup

```go title="db.go"
import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // ลงทะเบียน driver ชื่อ "pgx"
)

db, err := sql.Open("pgx", "postgres://app:secret@localhost:5432/app")
```

| Database | Driver import | ชื่อ |
| --- | --- | --- |
| PostgreSQL | `github.com/jackc/pgx/v5/stdlib` | `pgx` |
| MySQL | `github.com/go-sql-driver/mysql` | `mysql` |
| SQLite (Go ล้วน) | `modernc.org/sqlite` | `sqlite` |

## Type หลัก

- `*sql.DB`: **pool** ของ connection ใช้พร้อมกันหลาย goroutine ได้ปลอดภัย สร้างหนึ่งตัวต่อ database แล้วแชร์กัน
- `*sql.Rows`, `*sql.Row`: ผลลัพธ์ของ query
- `*sql.Tx`: transaction
- `*sql.Stmt`: prepared statement

:::caution[ข้อควรระวัง]
- อย่าเปิด `*sql.DB` ใหม่ทุก request ให้เปิดครั้งเดียวตอนเริ่มโปรแกรมแล้วส่งต่อเข้าไป
- placeholder ต่างกันตาม driver: `$1` สำหรับ Postgres และ `?` สำหรับ MySQL/SQLite
:::
