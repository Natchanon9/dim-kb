---
title: "Dependency Injection"
description: "ส่ง dependency เข้าไปผ่าน constructor"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

ใน Go DI หมายถึง: type รับสิ่งที่ต้องใช้เป็น parameter ของ constructor (ส่วนใหญ่เป็น interface) แล้ว `main` ต่อทุกอย่างเข้าด้วยกัน ไม่ต้องใช้ framework

## ตัวอย่าง

```go title="greeter.go"
type Clock interface{ Now() time.Time }

type Greeter struct{ clock Clock }

func NewGreeter(c Clock) *Greeter { return &Greeter{clock: c} }

func (g *Greeter) Greet() string {
	if g.clock.Now().Hour() < 12 {
		return "good morning"
	}
	return "good afternoon"
}

// ใน test: นาฬิกาที่ตั้งเวลาตายตัว → "good morning"
```

## ต่อใน main

```go
db := mustOpenDB(cfg.DatabaseURL)
repo := order.NewRepository(db)
svc := order.NewService(repo, clock.Real{})
handler := order.NewHandler(svc)
```

ตัว generate โค้ด (Wire) หรือ container (Fx) ช่วยได้เฉพาะในแอปที่ใหญ่มาก

:::caution[ข้อควรระวัง]
- เลี่ยงตัวแปร global ระดับ package สำหรับ dependency (`var DB *sql.DB`) เพราะซ่อนการผูกกัน และทำให้ test แชร์ state กัน
:::
