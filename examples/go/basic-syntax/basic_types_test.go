package basicsyntax

import "fmt"

func Example_basicTypes() {
	var i int = 42
	var f float64 = 3.14
	var b byte = 'A'
	var r rune = 'ก'
	var s string = "hi"
	fmt.Printf("%T %T %T %T %T\n", i, f, b, r, s)
	// Output: int float64 uint8 int32 string
}
