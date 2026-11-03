package functions

import "fmt"

func counter() func() int {
	n := 0
	return func() int {
		n++ // n lives on between calls
		return n
	}
}

func Example_closures() {
	next := counter()
	next()
	next()
	other := counter() // its own n
	fmt.Println(next(), other())
	// Output: 3 1
}
