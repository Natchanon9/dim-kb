package arraysslices

import "fmt"

func Example_arrays() {
	var a [3]int
	b := [...]string{"x", "y"} // length from the literal
	a[0] = 1
	c := a // copies all elements
	c[0] = 99
	fmt.Println(a, len(b), c, a == [3]int{1, 0, 0})
	// Output: [1 0 0] 2 [99 0 0] true
}
