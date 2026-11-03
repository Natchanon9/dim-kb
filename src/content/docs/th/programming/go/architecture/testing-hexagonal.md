---
title: "Test โค้ดแบบ Hexagonal"
description: "test core ด้วย mock และ test adapter กับ dependency จริง"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

port ทำให้ test ตรงไปตรงมา: unit test core ด้วย adapter ปลอม และ test adapter จริงแต่ละตัวแยกกับ dependency จริง

## Mock ที่เขียนเอง

```go title="service_test.go"
type mockRepo struct {
	account Account
	getErr  error
	saveErr error
	saved   []Account // จดการเรียก
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

## Test pyramid ของโครงแบบนี้

| ระดับ | อะไร | อย่างไร |
| --- | --- | --- |
| Unit | service ใน core | fake/mock ของ port |
| Adapter | repository, HTTP client | Postgres จริง (testcontainers) / `httptest` |
| Handler | route, status code | `app.Test` / `httptest` กับ service ปลอม |
| End-to-end | ทั้ง service | น้อยๆ ช้า รันใน CI |

interface ใหญ่ใช้ `go.uber.org/mock` (`mockgen`) หรือ `testify/mock` generate mock ให้

:::caution[ข้อควรระวัง]
- test ที่ใช้ mock หนักและ assert ทุกการเรียกจะพังทุกครั้งที่ refactor แม้ไม่ได้ทำอะไรผิด ให้ assert ผลลัพธ์ก่อน
:::
