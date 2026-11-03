package architecture

import (
	"context"
	"errors"
	"testing"
)

// mockRepo is a hand-written mock: canned answers plus a record of calls.
type mockRepo struct {
	account Account
	getErr  error
	saveErr error
	saved   []Account
}

func (m *mockRepo) Get(context.Context, string) (Account, error) { return m.account, m.getErr }
func (m *mockRepo) Save(_ context.Context, a Account) error {
	m.saved = append(m.saved, a)
	return m.saveErr
}

func TestWithdraw(t *testing.T) {
	dbDown := errors.New("db down")
	tests := []struct {
		name      string
		repo      *mockRepo
		amount    int
		wantErr   error
		wantSaved int
	}{
		{"ok", &mockRepo{account: Account{"a1", 100}}, 30, nil, 1},
		{"insufficient", &mockRepo{account: Account{"a1", 10}}, 30, ErrInsufficientFunds, 0},
		{"get fails", &mockRepo{getErr: dbDown}, 30, dbDown, 0},
		{"save fails", &mockRepo{account: Account{"a1", 100}, saveErr: dbDown}, 30, dbDown, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAccountService(tt.repo).Withdraw(context.Background(), "a1", tt.amount)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("err = %v, want %v", err, tt.wantErr)
			}
			if len(tt.repo.saved) != tt.wantSaved {
				t.Errorf("Save called %d times, want %d", len(tt.repo.saved), tt.wantSaved)
			}
		})
	}
}
