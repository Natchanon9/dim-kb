package goroutines

import (
	"fmt"
	"runtime"
	"strings"
)

func Example_debugging() {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, true) // true: dump every goroutine
	fmt.Println(strings.HasPrefix(string(buf[:n]), "goroutine "))
	// Output: true
}
