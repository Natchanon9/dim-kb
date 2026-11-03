---
title: "Hexagonal Architecture"
description: "port และ adapter: แยก business logic ออกจาก DB และ HTTP"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

hexagonal architecture (ports & adapters) วาง business logic ไว้ใน core ที่พึ่งแค่ interface (**port**) ส่วน database, HTTP และ queue เป็น **adapter** ที่ implement หรือเรียก port เหล่านั้น

```text
HTTP handler ──▶ [ AccountService ] ──▶ AccountRepository (port)
 (driving adapter)     core                     ▲
                                     Postgres / memory adapter
```

## Core และ port

```go title="core.go"
type AccountRepository interface { // port: สิ่งที่ core ต้องการ
	Get(ctx context.Context, id string) (Account, error)
	Save(ctx context.Context, a Account) error
}

type AccountService struct{ repo AccountRepository }

func (s *AccountService) Withdraw(ctx context.Context, id string, amount int) (int, error) {
	a, err := s.repo.Get(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("withdraw: %w", err)
	}
	if a.Balance < amount {
		return a.Balance, ErrInsufficientFunds
	}
	a.Balance -= amount
	return a.Balance, s.repo.Save(ctx, a)
}
```

## Adapter ต่อกันใน main

```go
repo := NewMemoryAccounts(Account{ID: "a1", Balance: 100}) // หรือ NewPostgresAccounts(db)
svc := NewAccountService(repo)
svc.Withdraw(ctx, "a1", 30)  // 70 <nil>
svc.Withdraw(ctx, "a1", 100) // 70 insufficient funds
```

core ไม่ import Fiber, GORM หรือ package ของ SQL เลย จึง test ง่าย และย้ายไป framework อื่นได้

:::caution[ข้อควรระวัง]
- อย่าให้ type ของ adapter หลุดเข้าไปใน core (model ของ GORM, `*fiber.Ctx`) ให้แปลงที่ขอบนอก
:::
