package moderngo

import "fmt"

func Example_loopvar() {
	var fns []func() int
	for i := range 3 {
		fns = append(fns, func() int { return i }) // Go 1.22+: each closure sees its own i
	}
	for _, f := range fns {
		fmt.Print(f())
	}
	fmt.Println()
	// Output: 012
}
