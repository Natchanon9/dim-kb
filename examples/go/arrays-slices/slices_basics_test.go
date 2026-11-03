package arraysslices

import "fmt"

func Example_slicesBasics() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s, len(s), s[1:3])
	// Output: [1 2 3 4] 4 [2 3]
}
