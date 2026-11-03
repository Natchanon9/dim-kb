package reflection

import (
	"fmt"
	"reflect"
)

func Example_basics() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v.Type(), v.Kind(), v.Float())
	// Output: float64 float64 3.4
}
