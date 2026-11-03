package arraysslices

import (
	"fmt"
	"slices"
)

func Example_slicesPackage() {
	s := []int{3, 1, 2}
	slices.Sort(s)
	i, found := slices.BinarySearch(s, 2)
	fmt.Println(s, slices.Contains(s, 3), slices.Index(s, 3), i, found, slices.Max(s))
	// Output: [1 2 3] true 2 1 true 3
}
