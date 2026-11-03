package functions

import "fmt"

func deferOrder() {
	for i := range 3 {
		defer fmt.Print(" ", i) // runs when deferOrder returns, last in first out
	}
	fmt.Print("start")
}

func deferArgs() {
	x := 1
	defer fmt.Println("deferred:", x) // x is evaluated now
	x = 2
	fmt.Println("now:", x)
}

func Example_defer() {
	deferOrder()
	fmt.Println()
	deferArgs()
	// Output:
	// start 2 1 0
	// now: 2
	// deferred: 1
}
