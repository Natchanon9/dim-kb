package architecture

import (
	"context"
	"fmt"
)

func Example_hexagonal() {
	repo := NewMemoryAccounts(Account{ID: "a1", Balance: 100}) // swap for a Postgres adapter in main
	svc := NewAccountService(repo)
	ctx := context.Background()

	fmt.Println(svc.Withdraw(ctx, "a1", 30))
	fmt.Println(svc.Withdraw(ctx, "a1", 100))
	fmt.Println(svc.Withdraw(ctx, "zz", 1))
	// Output:
	// 70 <nil>
	// 70 insufficient funds
	// 0 withdraw: account not found
}
