package reflection

import (
	"fmt"
	"reflect"
)

type Signup struct {
	Email string `validate:"required"`
	Nick  string
}

func missingRequired(v any) []string {
	val := reflect.ValueOf(v)
	t := val.Type()
	var missing []string
	for i := range t.NumField() {
		if t.Field(i).Tag.Get("validate") == "required" && val.Field(i).IsZero() {
			missing = append(missing, t.Field(i).Name)
		}
	}
	return missing
}

func Example_tags() {
	fmt.Println(missingRequired(Signup{Nick: "x"}))
	// Output: [Email]
}
