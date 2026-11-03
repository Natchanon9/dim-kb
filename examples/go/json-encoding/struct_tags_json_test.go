package jsonencoding

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price,omitempty"` // dropped when zero
	Stock    int     `json:"stock,string"`    // encoded as a JSON string
	Internal string  `json:"-"`               // never encoded
}

func Example_tags() {
	b, _ := json.Marshal(Product{ID: 1, Name: "pen", Stock: 5, Internal: "x"})
	fmt.Println(string(b))
	// Output: {"id":1,"name":"pen","stock":"5"}
}
