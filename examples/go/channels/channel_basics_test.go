package channels

import "fmt"

func Example_basics() {
	ch := make(chan string)
	go func() { ch <- "ping" }() // send
	msg := <-ch                  // receive: waits for the send
	fmt.Println(msg)
	// Output: ping
}
