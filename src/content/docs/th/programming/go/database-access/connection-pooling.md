---
title: "Connection Pooling"
description: "กำหนดขนาดและจูน pool ของ *sql.DB"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`*sql.DB` เก็บ pool ของ connection ค่าเริ่มต้นเปิดได้ไม่จำกัด ตอน load สูงอาจทำให้ database รับไม่ไหว ให้ตั้งขีดจำกัดเอง

## ตัวอย่าง

```go title="pool.go"
db.SetMaxOpenConns(10)                  // เพดาน
db.SetMaxIdleConns(5)                   // เก็บไว้พร้อมใช้ระหว่าง request
db.SetConnMaxLifetime(30 * time.Minute) // เปลี่ยน connection ที่อายุยาว
db.SetConnMaxIdleTime(5 * time.Minute)

fmt.Println(db.Stats().MaxOpenConnections, db.Stats().OpenConnections) // 10 0
```

## กำหนดขนาด

- ผลรวมของทุก instance ต้องน้อยกว่า `max_connections` ของ database (Postgres ค่าเริ่มต้น 100)
- เริ่มที่ 10–25 ต่อ instance เป็นค่าที่นิยม แล้วดู `db.Stats()` (`WaitCount`, `WaitDuration`) เพื่อจูน
- platform แบบ serverless ที่มี instance เยอะมักต้องใช้ proxy อย่าง PgBouncer

:::caution[ข้อควรระวัง]
- `rows` ที่ไม่ปิดและ transaction ที่ยาวจะถือ connection ไว้ อาการคือ request ค้างพร้อม `WaitCount` ที่สูงขึ้น
:::
