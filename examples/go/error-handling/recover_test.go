package errorhandling

import "fmt"

func safeDivide(a, b int) (q int, err error) {
	defer func() {
		if r := recover(); r != nil { // only works inside a deferred func
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

func Example_recover() {
	fmt.Println(safeDivide(6, 3))
	fmt.Println(safeDivide(1, 0))
	// Output:
	// 2 <nil>
	// 0 recovered: runtime error: integer divide by zero
}
