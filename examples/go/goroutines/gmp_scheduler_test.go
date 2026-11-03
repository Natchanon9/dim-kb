package goroutines

import (
	"fmt"
	"runtime"
)

// Compiled but not run: the numbers depend on the machine.
func Example_scheduler() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Ps (GOMAXPROCS):", runtime.GOMAXPROCS(0))
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
