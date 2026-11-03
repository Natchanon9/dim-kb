package errorhandling

import "fmt"

func mustPositive(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("mustPositive: got %d", n)) // a bug, not a normal error
	}
	return n
}

func Example_panic() {
	defer func() { fmt.Println("recovered:", recover()) }()
	mustPositive(-1)
	fmt.Println("never printed")
	// Output: recovered: mustPositive: got -1
}
