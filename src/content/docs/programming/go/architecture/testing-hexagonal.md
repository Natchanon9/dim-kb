---
title: "Testing Hexagonal Code"
description: "Test the core with mocks, adapters with real dependencies"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Ports make testing straightforward: unit-test the core with a fake adapter, and test each real adapter separately against the real dependency.

## Hand-written mock

```go title="service_test.go"
type mockRepo struct {
	account Account
	getErr  error
	saveErr error
	saved   []Account // records calls
}

func (m *mockRepo) Get(context.Context, string) (Account, error) { return m.account, m.getErr }
func (m *mockRepo) Save(_ context.Context, a Account) error {
	m.saved = append(m.saved, a)
	return m.saveErr
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name      string
		repo      *mockRepo
		wantErr   error
		wantSaved int
	}{
		{"ok", &mockRepo{account: Account{"a1", 100}}, nil, 1},
		{"insufficient", &mockRepo{account: Account{"a1", 10}}, ErrInsufficientFunds, 0},
		{"get fails", &mockRepo{getErr: dbDown}, dbDown, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAccountService(tt.repo).Withdraw(context.Background(), "a1", 30)
			if !errors.Is(err, tt.wantErr) { ... }
			if len(tt.repo.saved) != tt.wantSaved { ... }
		})
	}
}
```

## Test pyramid for this layout

| Level | What | How |
| --- | --- | --- |
| Unit | Core services | Fakes/mocks of ports |
| Adapter | Repository, HTTP client | Real Postgres (testcontainers) / `httptest` |
| Handler | Routes, status codes | `app.Test` / `httptest` with a fake service |
| End-to-end | Whole service | Few, slow, in CI |

For big interfaces, `go.uber.org/mock` (`mockgen`) or `testify/mock` generate mocks.

:::caution[Pitfalls]
- Mock-heavy tests that assert every call break on harmless refactors; assert outcomes first.
:::
