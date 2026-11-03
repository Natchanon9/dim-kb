package pointers

import "fmt"

func Example_newVsMake() {
	p := new(int) // *int pointing to a zero int
	*p = 7
	s := make([]int, 2)       // ready-to-use slice
	m := make(map[string]int) // ready-to-use map
	ch := make(chan int, 1)   // ready-to-use channel
	m["a"] = 1
	ch <- 9
	fmt.Println(*p, s, m, <-ch)

	ps := new([]int) // pointer to a nil slice: rarely what you want
	fmt.Println(*ps == nil)
	// Output:
	// 7 [0 0] map[a:1] 9
	// true
}
