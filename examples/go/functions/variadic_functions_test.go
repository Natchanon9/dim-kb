package functions

import "fmt"

func sum(nums ...int) int { // nums is a []int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Example_variadic() {
	fmt.Println(sum(), sum(1, 2, 3))
	xs := []int{4, 5}
	fmt.Println(sum(xs...)) // spread a slice
	// Output:
	// 0 6
	// 9
}
