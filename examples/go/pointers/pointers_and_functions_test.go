package pointers

import "fmt"

func reset(n *int) { *n = 0 }

func newInt(v int) *int { return &v } // safe: v moves to the heap

func Example_functions() {
	x := 5
	reset(&x)
	p := newInt(3)
	fmt.Println(x, *p)
	// Output: 0 3
}
