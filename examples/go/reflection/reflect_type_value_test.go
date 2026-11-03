package reflection

import (
	"fmt"
	"reflect"
)

type Celsius float64

func Example_typeValue() {
	t := reflect.TypeOf(Celsius(1))
	fmt.Println(t.Name(), t.Kind()) // Type: the declared type; Kind: the underlying category

	x := 1
	v := reflect.ValueOf(&x).Elem() // must go through a pointer to modify
	fmt.Println(v.CanSet())
	v.SetInt(42)
	fmt.Println(x)
	// Output:
	// Celsius float64
	// true
	// 42
}
