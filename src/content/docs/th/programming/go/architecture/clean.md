---
title: "Clean Architecture"
description: "entity, use case และ adapter ที่ dependency ชี้เข้าข้างในเสมอ"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Clean Architecture จัดโค้ดเป็นวงซ้อน: **entity** (กฎทางธุรกิจ), **use case** (การกระทำของแอป), **interface adapter** (controller, repository), **framework** (DB, web) โดย dependency ชี้เข้าข้างในอย่างเดียว

## Use case กับ interface ของตัวเอง

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
	u, err := NewUser(in.Email, hash) // entity ตรวจ email
	if err != nil {
		return RegisterOutput{}, err
	}
	return RegisterOutput{Email: u.Email}, uc.Users.Create(ctx, u)
}
// "Ann@Example.com" → {ann@example.com} ; ซ้ำอีกรอบ → email already registered
```

## Hexagonal vs Clean

แนวคิดหลักเหมือนกัน: business logic พึ่ง interface ไม่ใช่ infrastructure ส่วน Clean เพิ่มชั้นให้ชัดขึ้น (entity กับ use case, model ขาเข้า/ขาออก) ใน Go ทำให้เบาเข้าไว้: หนึ่ง package ต่อหนึ่ง feature ที่มี handler, service และ repository มักพอแล้ว

:::caution[ข้อควรระวัง]
- 4 ชั้นกับ mapper ทุก field สำหรับ endpoint CRUD ธรรมดาเกินจำเป็น เพิ่มโครงสร้างเมื่อ logic โตขึ้นจริง
:::
