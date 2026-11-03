package functions

import (
	"fmt"
	"strconv"
)

type celsius float64

func (c celsius) String() string { return fmt.Sprintf("%.1f°C", float64(c)) }

var _ fmt.Stringer = celsius(0) // compile-time check: celsius implements Stringer

func Example_blank() {
	n, _ := strconv.Atoi("42") // ignore the error on purpose
	for _, v := range []celsius{21.5} {
		fmt.Println(n, v)
	}
	// Output: 42 21.5°C
}
