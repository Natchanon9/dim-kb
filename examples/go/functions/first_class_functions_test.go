package functions

import "fmt"

func apply(nums []int, f func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

func Example_firstClass() {
	square := func(n int) int { return n * n }
	fmt.Println(apply([]int{1, 2, 3}, square))

	ops := map[string]func(a, b int) int{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
	}
	fmt.Println(ops["mul"](3, 4))
	// Output:
	// [1 4 9]
	// 12
}
