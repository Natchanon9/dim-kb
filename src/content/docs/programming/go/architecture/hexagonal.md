---
title: "Hexagonal Architecture"
description: "Ports and adapters: keep business logic independent of DB and HTTP"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

Hexagonal architecture (ports & adapters) puts business logic in a core that depends only on interfaces (**ports**). Databases, HTTP and queues are **adapters** that implement or call those ports.

```text
HTTP handler ──▶ [ AccountService ] ──▶ AccountRepository (port)
 (driving adapter)     core                     ▲
                                     Postgres / memory adapter
```

## Core and port

```go title="core.go"
type AccountRepository interface { // port: what the core needs
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

## Adapters, wired in main

```go
repo := NewMemoryAccounts(Account{ID: "a1", Balance: 100}) // or NewPostgresAccounts(db)
svc := NewAccountService(repo)
svc.Withdraw(ctx, "a1", 30)  // 70 <nil>
svc.Withdraw(ctx, "a1", 100) // 70 insufficient funds
```

The core imports no Fiber, GORM or SQL packages, so it is easy to test and to move to another framework.

:::caution[Pitfalls]
- Don't leak adapter types into the core (GORM models, `*fiber.Ctx`); map them at the edge.
:::
