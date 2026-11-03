package basicsyntax

import (
	"fmt"
	"strconv"
)

func Example_typeConversions() {
	i := 42
	f := float64(i)
	x := 3.9
	fmt.Println(f, int(x)) // int() truncates toward zero

	s := strconv.Itoa(i)
	n, err := strconv.Atoi("123")
	fmt.Println(s, n, err, string(rune(65)))
	// Output:
	// 42 3
	// 42 123 <nil> A
}
