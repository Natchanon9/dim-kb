package gomaps

import "fmt"

func Example_mapBasics() {
	ages := map[string]int{"ann": 30}
	ages["bob"] = 25
	fmt.Println(ages["ann"], ages["zed"], len(ages)) // missing key → zero value

	var m map[string]int // nil map: reads work, writes panic
	fmt.Println(m["x"], m == nil)
	// Output:
	// 30 0 2
	// 0 true
}
