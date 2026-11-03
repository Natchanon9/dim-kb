---
title: "Clean Architecture"
description: "Entities, use cases and adapters with dependencies pointing inward"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Clean Architecture arranges code in rings: **entities** (business rules), **use cases** (application actions), **interface adapters** (controllers, repositories), **frameworks** (DB, web). Dependencies only point inward.

## Use case with its own interfaces

```go title="register.go"
type UserStore interface {
	Exists(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, u User) error
}

type PasswordHasher interface{ Hash(pw string) (string, error) }

type RegisterUser struct {
	Users  UserStore
	Hasher PasswordHasher
}

func (uc RegisterUser) Execute(ctx context.Context, in RegisterInput) (RegisterOutput, error) {
	taken, err := uc.Users.Exists(ctx, strings.ToLower(in.Email))
	if err != nil {
		return RegisterOutput{}, err
	}
	if taken {
		return RegisterOutput{}, ErrEmailTaken
	}
	hash, _ := uc.Hasher.Hash(in.Password)
	u, err := NewUser(in.Email, hash) // entity validates the email
	if err != nil {
		return RegisterOutput{}, err
	}
	return RegisterOutput{Email: u.Email}, uc.Users.Create(ctx, u)
}
// "Ann@Example.com" → {ann@example.com} ; again → email already registered
```

## Hexagonal vs Clean

Same core idea: business logic depends on interfaces, not on infrastructure. Clean adds explicit layers (entities vs use cases, input/output models). In Go, keep it light: one package per feature with handler, service and repository is often enough.

:::caution[Pitfalls]
- Four layers and a mapper per field for a CRUD endpoint is overkill; add structure when the logic grows.
:::
