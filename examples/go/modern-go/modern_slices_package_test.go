//go:build go1.23

package moderngo

import (
	"fmt"
	"slices"
)

func Example_modernSlices() {
	s := []int{3, 1, 2}
	fmt.Println(slices.Sorted(slices.Values(s))) // iterator → sorted slice
	fmt.Println(slices.Collect(Countdown(3)))    // any iterator → slice
	for chunk := range slices.Chunk([]int{1, 2, 3, 4, 5}, 2) {
		fmt.Println(chunk)
	}
	// Output:
	// [1 2 3]
	// [3 2 1]
	// [1 2]
	// [3 4]
	// [5]
}
