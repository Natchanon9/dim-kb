package generics

import (
	"fmt"
	"io"
	"strings"
)

// An interface is enough here: the function only calls methods.
func countBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}

// Generics fit when the same algorithm needs the element type itself.
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}

func Example_bestPractices() {
	n, _ := countBytes(strings.NewReader("hello"))
	v, ok := Last([]int{1, 2, 3})
	fmt.Println(n, v, ok)
	// Output: 5 3 true
}
