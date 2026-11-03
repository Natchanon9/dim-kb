package pointers

import "fmt"

func setNil(pp **int) { *pp = nil } // change the caller's pointer itself

func Example_pointerToPointer() {
	x := 1
	p := &x
	pp := &p
	**pp = 2
	fmt.Println(x, *pp == p)
	setNil(&p)
	fmt.Println(p == nil)
	// Output:
	// 2 true
	// true
}
