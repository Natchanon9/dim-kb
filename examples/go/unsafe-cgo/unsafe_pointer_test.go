package unsafecgo

import (
	"fmt"
	"unsafe"
)

func Example_unsafePointer() {
	b := []byte("hello")
	s := unsafe.String(unsafe.SliceData(b), len(b)) // Go 1.20+: no copy; b must never change again
	fmt.Println(s)

	x := int32(7)
	p := unsafe.Pointer(&x) // any pointer → unsafe.Pointer → another pointer type
	fmt.Println(*(*int32)(p))
	// Output:
	// hello
	// 7
}
