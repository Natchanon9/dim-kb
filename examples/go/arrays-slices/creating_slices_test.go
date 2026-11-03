package arraysslices

import "fmt"

func Example_creatingSlices() {
	a := []int{1, 2}        // literal
	b := make([]int, 3)     // len 3, all zero
	c := make([]int, 0, 10) // len 0, room for 10
	var d []int             // nil slice
	fmt.Println(a, b, len(c), cap(c), d == nil)
	// Output: [1 2] [0 0 0] 0 10 true
}
