package architecture

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// ---- core: domain ----

type Account struct {
	ID      string
	Balance int
}

var (
	ErrNotFound          = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// ---- port: what the core needs from the outside ----

type AccountRepository interface {
	Get(ctx context.Context, id string) (Account, error)
	Save(ctx context.Context, a Account) error
}

// ---- core: application service ----

type AccountService struct{ repo AccountRepository }

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Withdraw(ctx context.Context, id string, amount int) (int, error) {
	a, err := s.repo.Get(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("withdraw: %w", err)
	}
	if a.Balance < amount {
		return a.Balance, ErrInsufficientFunds
	}
	a.Balance -= amount
	if err := s.repo.Save(ctx, a); err != nil {
		return 0, fmt.Errorf("withdraw: %w", err)
	}
	return a.Balance, nil
}

// ---- adapter: an in-memory implementation of the port (Postgres would be another) ----

type MemoryAccounts struct {
	mu   sync.Mutex
	data map[string]Account
}

func NewMemoryAccounts(accounts ...Account) *MemoryAccounts {
	m := &MemoryAccounts{data: map[string]Account{}}
	for _, a := range accounts {
		m.data[a.ID] = a
	}
	return m
}

func (m *MemoryAccounts) Get(_ context.Context, id string) (Account, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	a, ok := m.data[id]
	if !ok {
		return Account{}, ErrNotFound
	}
	return a, nil
}

func (m *MemoryAccounts) Save(_ context.Context, a Account) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[a.ID] = a
	return nil
}
