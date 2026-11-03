package basicsyntax

import "fmt"

// MaxItems is exported: it starts with a capital letter.
const MaxItems = 10

// discount is unexported: only code in this package can call it.
func discount(price float64) float64 { return price * 0.9 }

func Example_naming() {
	fmt.Println(MaxItems, discount(100))
	// Output: 10 90
}
