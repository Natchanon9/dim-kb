package structs

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID    int      `json:"id"`
	Items []string `json:"items"`
	Note  string   `json:"note,omitempty"`
}

func Example_json() {
	data, _ := json.Marshal(Order{ID: 1, Items: []string{"pen"}})
	fmt.Println(string(data))

	var o Order
	err := json.Unmarshal([]byte(`{"id":2,"items":["a","b"],"extra":true}`), &o) // unknown fields ignored
	fmt.Println(o.ID, o.Items, err)
	// Output:
	// {"id":1,"items":["pen"]}
	// 2 [a b] <nil>
}
