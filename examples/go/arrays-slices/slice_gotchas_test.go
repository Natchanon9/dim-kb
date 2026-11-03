package arraysslices

import "fmt"

func Example_sliceGotchas() {
	a := make([]int, 3, 4)
	b := append(a, 1) // fits in cap: writes a's backing array
	c := append(a, 2) // same slot again
	fmt.Println(b[3], c[3])

	safe := append(a[:3:3], 1) // full slice expression caps it, forcing a copy
	fmt.Println(safe[3], b[3])
	// Output:
	// 2 2
	// 1 2
}
