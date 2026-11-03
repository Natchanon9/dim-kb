package generics

import (
	"cmp"
	"fmt"
)

func Larger[T cmp.Ordered](a, b T) T { // T must support < and >
	if a > b {
		return a
	}
	return b
}

func Example_typeParameters() {
	fmt.Println(Larger(3, 7), Larger("a", "b"), Larger[float64](1, 2.5))
	// Output: 7 b 2.5
}
