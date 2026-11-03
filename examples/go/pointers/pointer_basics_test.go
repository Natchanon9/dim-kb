package pointers

import "fmt"

func Example_basics() {
	x := 10
	p := &x // p is a *int holding x's address
	*p = 20 // write through the pointer
	fmt.Println(x, *p, p != nil)
	// Output: 20 20 true
}
