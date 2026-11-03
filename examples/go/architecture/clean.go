package architecture

import (
	"context"
	"errors"
	"strings"
)

// ---- entities: business rules that don't depend on anything ----

type User struct {
	Email        string
	PasswordHash string
}

func NewUser(email, hash string) (User, error) {
	if !strings.Contains(email, "@") {
		return User{}, errors.New("invalid email")
	}
	return User{Email: strings.ToLower(email), PasswordHash: hash}, nil
}

// ---- use case layer: one application action, with the interfaces it needs ----

var ErrEmailTaken = errors.New("email already registered")

type UserStore interface {
	Exists(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, u User) error
}

type PasswordHasher interface {
	Hash(pw string) (string, error)
}

type RegisterInput struct{ Email, Password string }
type RegisterOutput struct{ Email string }

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
	hash, err := uc.Hasher.Hash(in.Password)
	if err != nil {
		return RegisterOutput{}, err
	}
	u, err := NewUser(in.Email, hash)
	if err != nil {
		return RegisterOutput{}, err
	}
	if err := uc.Users.Create(ctx, u); err != nil {
		return RegisterOutput{}, err
	}
	return RegisterOutput{Email: u.Email}, nil
}
