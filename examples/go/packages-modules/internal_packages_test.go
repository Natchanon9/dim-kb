package packagesmodules

import (
	"fmt"

	"github.com/Natchanon9/dim-kb/examples/go/packages-modules/internal/format"
)

func Example_internal() {
	// allowed: this package lives under packages-modules/, the parent of internal/
	fmt.Println(format.Title("internal"))
	// Output: Internal
}
