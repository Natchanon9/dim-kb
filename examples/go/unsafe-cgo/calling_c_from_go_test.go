//go:build cgo

package unsafecgo

import "fmt"

func Example_callingC() {
	fmt.Println(CStrlen("hello, C"))
	// Output: 8
}
