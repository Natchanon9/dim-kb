//go:build go1.23

package moderngo

import (
	"fmt"
	"iter"
)

func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func Example_iterators() {
	for i, s := range Enumerate([]string{"a", "b"}) {
		fmt.Println(i, s)
	}

	next, stop := iter.Pull(Countdown(3)) // turn a push iterator into next()
	defer stop()
	v1, _ := next()
	v2, _ := next()
	fmt.Println(v1, v2)
	// Output:
	// 0 a
	// 1 b
	// 3 2
}
