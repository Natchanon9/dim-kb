package bestpractices

import "fmt"

// Greet returns a greeting for name.
//
// It is safe for concurrent use. See also [Greeter.Greet].
func Greet(name string) string { return "Hello, " + name }

// ExampleGreet appears under Greet on pkg.go.dev and runs as a test.
func ExampleGreet() {
	fmt.Println(Greet("docs"))
	// Output: Hello, docs
}
