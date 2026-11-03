package channels

import "fmt"

func Example_buffered() {
	ch := make(chan int, 2)
	ch <- 1 // no receiver needed while there is room
	ch <- 2
	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch, <-ch) // FIFO
	// Output:
	// 2 2
	// 1 2
}
