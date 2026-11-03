package arraysslices

import "fmt"

func Example_appendCopy() {
	s := make([]int, 0, 2)
	s = append(s, 1, 2)
	before := cap(s)
	s = append(s, 3) // over capacity: new, bigger backing array
	fmt.Println(s, before, cap(s) > before)

	dst := make([]int, 2)
	n := copy(dst, []int{7, 8, 9}) // copies min(len(dst), len(src))
	fmt.Println(dst, n)
	// Output:
	// [1 2 3] 2 true
	// [7 8] 2
}
