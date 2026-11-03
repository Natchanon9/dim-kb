package packagesmodules

import (
	"fmt"
	str "strings" // alias

	g "github.com/Natchanon9/dim-kb/examples/go/packages-modules/greet"
)

func Example_importing() {
	fmt.Println(str.Repeat("go", 2), g.Hello("alias"))
	// Output: gogo Hello, alias
}
