package jsonencoding

import (
	"encoding/json"
	"fmt"
)

func Example_basics() {
	b, _ := json.Marshal(map[string]any{"name": "go", "tags": []string{"a"}, "ok": true})
	fmt.Println(string(b)) // map keys come out sorted

	var v map[string]any
	json.Unmarshal([]byte(`{"n":1}`), &v)
	fmt.Printf("%T\n", v["n"]) // numbers in `any` become float64

	pretty, _ := json.MarshalIndent(map[string]int{"a": 1}, "", "  ")
	fmt.Println(string(pretty))
	// Output:
	// {"name":"go","ok":true,"tags":["a"]}
	// float64
	// {
	//   "a": 1
	// }
}
