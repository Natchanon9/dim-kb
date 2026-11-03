package goroutines

import "fmt"

func Example_lifecycle() {
	done := make(chan struct{})
	go func() {
		defer close(done) // signal: this goroutine has finished
		fmt.Println("working")
	}()
	<-done // without this, main could return first and kill the goroutine
	fmt.Println("finished")
	// Output:
	// working
	// finished
}
