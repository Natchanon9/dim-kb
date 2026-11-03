package arraysslices

import "fmt"

func Example_sliceInternals() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] // pointer to arr[1], len 2, cap 4
	fmt.Println(s, len(s), cap(s))
	s[0] = 20 // writes into arr
	fmt.Println(arr)
	// Output:
	// [2 3] 2 4
	// [1 20 3 4 5]
}
