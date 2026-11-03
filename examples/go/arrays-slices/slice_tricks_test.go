package arraysslices

import (
	"fmt"
	"slices"
)

func Example_sliceTricks() {
	s := []int{1, 2, 3, 4, 5}
	s = slices.Delete(s, 1, 2) // remove index 1
	fmt.Println(s)
	s = slices.Insert(s, 1, 2) // put 2 back at index 1
	fmt.Println(s)
	slices.Reverse(s)
	fmt.Println(s)

	evens := s[:0] // filter in place, reusing the backing array
	for _, v := range s {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	fmt.Println(evens)
	// Output:
	// [1 3 4 5]
	// [1 2 3 4 5]
	// [5 4 3 2 1]
	// [4 2]
}
