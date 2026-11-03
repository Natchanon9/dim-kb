package controlflow

import (
	"fmt"
	"strconv"
)

func Example_ifWithInit() {
	stock := map[string]int{"apple": 3}
	if n, ok := stock["apple"]; ok {
		fmt.Println("in stock:", n)
	}
	if _, err := strconv.Atoi("x"); err != nil {
		fmt.Println("error:", err)
	}
	// n and err don't exist here
	// Output:
	// in stock: 3
	// error: strconv.Atoi: parsing "x": invalid syntax
}
