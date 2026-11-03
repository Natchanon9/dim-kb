//go:build cgo

package unsafecgo

import "fmt"

func Example_cgoBasics() {
	fmt.Println(Add(2, 3))
	// Output: 5
}
