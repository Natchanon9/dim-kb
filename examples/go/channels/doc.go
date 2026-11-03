// Package channels holds the snippets for the "Channels" pages
// in src/content/docs/programming/go/channels/.
package channels

// emit returns a channel that yields vals and then closes.
func emit(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}
