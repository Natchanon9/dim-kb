package packagesmodules

import (
	"fmt"

	"github.com/Natchanon9/dim-kb/examples/go/packages-modules/greet"
)

func Example_packages() {
	fmt.Println(greet.Hello("gopher"))
	fmt.Println(greet.Excited("go"))
	// greet.suffix() would not compile: unexported
	// Output:
	// Hello, gopher
	// Hello, go!
}
