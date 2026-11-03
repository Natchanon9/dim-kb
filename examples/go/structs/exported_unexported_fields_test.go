package structs

import (
	"encoding/json"
	"fmt"
)

type account struct {
	ID       int    // exported: visible to other packages and to encoding/json
	password string // unexported: package-private
}

func Example_exportedFields() {
	b, _ := json.Marshal(account{ID: 1, password: "x"})
	fmt.Println(string(b))
	// Output: {"ID":1}
}
