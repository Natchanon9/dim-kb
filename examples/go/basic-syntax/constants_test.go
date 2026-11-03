package basicsyntax

import (
	"fmt"
	"time"
)

const Pi = 3.14159 // untyped constant

func Example_constants() {
	const timeout = 5 * time.Second
	var f32 float32 = Pi // untyped Pi fits any float type
	fmt.Println(Pi, f32, timeout)
	// Output: 3.14159 3.14159 5s
}
