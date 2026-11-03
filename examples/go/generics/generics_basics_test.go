package generics

import (
	"fmt"
	"strconv"
)

// Map works for any element types T and U.
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func Example_basics() {
	doubled := Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) })
	fmt.Printf("%q\n", doubled)
	// Output: ["2" "4" "6"]
}
