package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Item struct {
	Name   string `json:"name"`
	Price  int    `json:"price,omitempty"`
	Secret string `json:"-"`
}

func Example_tags() {
	b, _ := json.Marshal(Item{Name: "pen", Secret: "x"})
	f, _ := reflect.TypeOf(Item{}).FieldByName("Name")
	fmt.Println(string(b), f.Tag.Get("json"))
	// Output: {"name":"pen"} name
}
