package architecture

import (
	"context"
	"fmt"
)

// interface adapters: in-memory store and a fake hasher
type memUsers map[string]User

func (m memUsers) Exists(_ context.Context, email string) (bool, error) {
	_, ok := m[email]
	return ok, nil
}
func (m memUsers) Create(_ context.Context, u User) error { m[u.Email] = u; return nil }

type reverseHasher struct{} // stand-in for bcrypt in tests

func (reverseHasher) Hash(pw string) (string, error) { return "hashed:" + pw, nil }

func Example_clean() {
	uc := RegisterUser{Users: memUsers{}, Hasher: reverseHasher{}}
	ctx := context.Background()

	fmt.Println(uc.Execute(ctx, RegisterInput{Email: "Ann@Example.com", Password: "pw"}))
	fmt.Println(uc.Execute(ctx, RegisterInput{Email: "ann@example.com", Password: "x"}))
	fmt.Println(uc.Execute(ctx, RegisterInput{Email: "nope", Password: "x"}))
	// Output:
	// {ann@example.com} <nil>
	// {} email already registered
	// {} invalid email
}
