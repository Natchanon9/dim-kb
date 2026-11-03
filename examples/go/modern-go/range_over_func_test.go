//go:build go1.23

package moderngo

import (
	"fmt"
	"iter"
)

// Countdown is an iterator: a func that calls yield for each value.
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i > 0; i-- {
			if !yield(i) { // false: the loop body hit break or return
				return
			}
		}
	}
}

func Example_rangeOverFunc() {
	var got []int
	for v := range Countdown(5) {
		if v == 2 {
			break
		}
		got = append(got, v)
	}
	fmt.Println(got)
	// Output: [5 4 3]
}
