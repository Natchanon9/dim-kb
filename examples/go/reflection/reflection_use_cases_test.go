package reflection

import (
	"fmt"
	"reflect"
)

// toMap is a tiny version of what encoders do: walk fields and read tags.
func toMap(v any) map[string]any {
	val := reflect.ValueOf(v)
	t := val.Type()
	out := map[string]any{}
	for i := range t.NumField() {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name
		}
		out[key] = val.Field(i).Interface()
	}
	return out
}

func Example_useCases() {
	fmt.Println(toMap(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{1, "go"}))
	// Output: map[id:1 name:go]
}
