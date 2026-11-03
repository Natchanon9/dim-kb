package structs

import "fmt"

func Example_literals() {
	a := User{Name: "Bob", Age: 20} // named fields: preferred
	b := User{"Cat", 30}            // positional: every field, in order
	p := &User{Name: "Dan"}         // pointer to a new struct
	p.Age = 40                      // same as (*p).Age
	fmt.Println(a, b, *p)
	// Output: {Bob 20} {Cat 30} {Dan 40}
}
