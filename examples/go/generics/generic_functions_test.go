package generics

import "fmt"

func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Example_functions() {
	evens := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	total := Reduce(evens, 0, func(acc, n int) int { return acc + n })
	fmt.Println(evens, total)
	// Output: [2 4] 6
}
