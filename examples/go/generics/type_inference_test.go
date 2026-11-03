package generics

import (
	"fmt"
	"strconv"
)

func Zero[T any]() T {
	var z T
	return z
}

func Example_inference() {
	fmt.Println(Map([]int{1, 2}, strconv.Itoa))    // T=int, U=string inferred
	fmt.Println(Zero[string]() == "", Zero[int]()) // no arguments: must be explicit
	// Output:
	// [1 2]
	// true 0
}
