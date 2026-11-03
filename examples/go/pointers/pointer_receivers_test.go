package pointers

import "fmt"

type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

func Example_pointerReceivers() {
	acc := Account{}
	acc.Deposit(50) // Go passes &acc automatically
	ptr := &acc
	ptr.Deposit(25)
	fmt.Println(acc.Balance)
	// Output: 75
}
