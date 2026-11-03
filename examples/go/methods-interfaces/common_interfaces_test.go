package methodsinterfaces

import (
	"fmt"
	"io"
	"strings"
)

type Temp float64

func (t Temp) String() string { return fmt.Sprintf("%.1f°C", float64(t)) } // fmt.Stringer

func Example_commonInterfaces() {
	fmt.Println(Temp(21.5))

	r := strings.NewReader("hello") // io.Reader
	var sb strings.Builder          // io.Writer
	n, _ := io.Copy(&sb, r)
	fmt.Println(n, sb.String())
	// Output:
	// 21.5°C
	// 5 hello
}
