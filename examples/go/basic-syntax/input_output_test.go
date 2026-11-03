package basicsyntax

import (
	"fmt"
	"strings"
)

type user struct {
	Name string
	Age  int
}

func Example_inputOutput() {
	u := user{"Ann", 30}
	fmt.Printf("%v | %+v | %T\n", u, u, u)
	fmt.Printf("[%6.2f] [%-4s] %q\n", 3.14159, "go", "hi")

	var name string
	var age int
	fmt.Fscan(strings.NewReader("Bob 25"), &name, &age) // fmt.Scan reads os.Stdin the same way
	fmt.Println(name, age)
	// Output:
	// {Ann 30} | {Name:Ann Age:30} | basicsyntax.user
	// [  3.14] [go  ] "hi"
	// Bob 25
}
