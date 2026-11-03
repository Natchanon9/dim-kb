package functions

import "fmt"

func add(a, b int) int { return a + b }

func double(n int) { n *= 2 } // n is a copy

func Example_basics() {
	x := 5
	double(x)
	fmt.Println(add(2, 3), x)
	// Output: 5 5
}
