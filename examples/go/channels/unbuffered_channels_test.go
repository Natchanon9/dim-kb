package channels

import "fmt"

func Example_unbuffered() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		v := <-ch
		fmt.Println("got", v)
		close(done)
	}()
	ch <- 42 // blocks until the goroutine receives: a handoff
	<-done
	// Output: got 42
}
