package functions

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func Example_multipleReturns() {
	q, err := divide(10, 4)
	fmt.Println(q, err)
	_, err = divide(1, 0)
	fmt.Println(err)
	// Output:
	// 2.5 <nil>
	// divide by zero
}
