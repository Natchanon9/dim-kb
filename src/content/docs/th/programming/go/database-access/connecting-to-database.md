---
title: "เชื่อมต่อ Database"
description: "เปิด ping และตั้งค่า pool ตอนเริ่มโปรแกรม"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`sql.Open` แค่ตรวจ argument ส่วนการเชื่อมต่อจริงเกิดตอนใช้ครั้งแรก ให้ ping ตอนเริ่มโปรแกรม DSN ที่ผิดจะได้พังเร็ว

## ตัวอย่าง

```go title="open.go"
func openDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}
```

## DSN

```text
postgres://user:password@host:5432/dbname?sslmode=disable
```

อ่านจาก environment variable (`DATABASE_URL`) อย่า hard-code รหัสผ่านไว้ในโค้ด ในเครื่องรัน Postgres ด้วย Docker:

```sh
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:17
```

:::caution[ข้อควรระวัง]
- `defer db.Close()` ควรอยู่ใน `main` ไม่ใช่ใน function ที่คืน db ออกไป
:::
