package jsonencoding

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"` // kept as raw bytes, decoded later
}

func Example_rawMessage() {
	var e Event
	json.Unmarshal([]byte(`{"type":"user","data":{"name":"ann"}}`), &e)
	switch e.Type {
	case "user":
		var u struct {
			Name string `json:"name"`
		}
		json.Unmarshal(e.Data, &u)
		fmt.Println(u.Name)
	}
	fmt.Println(string(e.Data))
	// Output:
	// ann
	// {"name":"ann"}
}
