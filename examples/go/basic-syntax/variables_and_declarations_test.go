package basicsyntax

import "fmt"

func Example_varDeclarations() {
	var name string = "gopher"
	var age = 3 // type inferred as int
	var (
		host = "localhost"
		port = 8080
	)
	fmt.Println(name, age, host, port)
	// Output: gopher 3 localhost 8080
}
