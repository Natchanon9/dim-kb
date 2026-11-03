package bestpractices

import (
	"bytes"
	"fmt"
	"runtime/pprof"
)

func busyWork() int {
	sum := 0
	for i := range 2_000_000 {
		sum += i % 7
	}
	return sum
}

func Example_pprof() {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil { // a file in real code: go tool pprof cpu.out
		panic(err)
	}
	busyWork()
	pprof.StopCPUProfile()
	fmt.Println(buf.Len() > 0)
	// Output: true
}
