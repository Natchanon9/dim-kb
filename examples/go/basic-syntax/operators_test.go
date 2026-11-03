package basicsyntax

import "fmt"

func Example_operators() {
	fmt.Println(7/2, 7%2, 7.0/2)
	fmt.Println(1<<3, 6&3, 6|3, 6^3, 6&^3)
	x := 5
	x += 2
	x++ // a statement, not an expression
	fmt.Println(x)
	// Output:
	// 3 1 3.5
	// 8 2 7 5 4
	// 8
}
