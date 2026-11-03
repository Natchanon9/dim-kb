package channels

import (
	"fmt"
	"time"
)

func Example_select() {
	fast := make(chan string, 1)
	fast <- "cache"
	select {
	case v := <-fast:
		fmt.Println(v)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}

	empty := make(chan int)
	select {
	case v := <-empty:
		fmt.Println(v)
	default: // runs immediately when no case is ready
		fmt.Println("no value")
	}
	// Output:
	// cache
	// no value
}
