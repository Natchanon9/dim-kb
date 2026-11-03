package reflection

import (
	"fmt"
	"reflect"
)

type Calc struct{}

func (Calc) Add(a, b int) int { return a + b }

func Example_dynamicCall() {
	m := reflect.ValueOf(Calc{}).MethodByName("Add") // look up by name at runtime
	out := m.Call([]reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)})
	fmt.Println(out[0].Int())
	// Output: 5
}
