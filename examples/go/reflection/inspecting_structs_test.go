package reflection

import (
	"fmt"
	"reflect"
)

type Account struct {
	Name  string
	Age   int
	email string
}

func Example_inspectStructs() {
	t := reflect.TypeOf(Account{})
	for i := range t.NumField() {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.IsExported())
	}
	// Output:
	// Name string true
	// Age int true
	// email string false
}
