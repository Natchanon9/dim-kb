package gomaps

import "fmt"

func Example_commaOk() {
	stock := map[string]int{"apple": 0}
	v, ok := stock["apple"] // present, value is 0
	fmt.Println(v, ok)
	v, ok = stock["pear"] // absent
	fmt.Println(v, ok)
	// Output:
	// 0 true
	// 0 false
}
