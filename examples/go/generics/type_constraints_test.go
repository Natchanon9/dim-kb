package generics

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n // allowed: every type in Number supports +
	}
	return total
}

type Score int

func Example_constraints() {
	fmt.Println(Sum([]int{1, 2}), Sum([]float64{1.5, 2}), Sum([]Score{3, 4}))
	// Output: 3 3.5 7
}
