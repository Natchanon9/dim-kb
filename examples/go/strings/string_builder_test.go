package gostrings

import (
	"fmt"
	"strings"
)

func Example_builder() {
	var sb strings.Builder // zero value is ready to use
	for i := range 3 {
		fmt.Fprintf(&sb, "%d,", i)
	}
	sb.WriteString("done")
	fmt.Println(sb.String(), sb.Len())
	// Output: 0,1,2,done 10
}
