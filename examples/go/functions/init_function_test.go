package functions

import "fmt"

var config map[string]string

func init() { // runs once, before main and before tests
	config = map[string]string{"env": "dev"}
}

func Example_init() {
	fmt.Println(config["env"])
	// Output: dev
}
