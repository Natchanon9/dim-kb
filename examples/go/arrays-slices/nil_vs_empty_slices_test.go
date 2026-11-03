package arraysslices

import (
	"encoding/json"
	"fmt"
)

func Example_nilVsEmpty() {
	var a []int  // nil
	b := []int{} // empty, not nil
	fmt.Println(a == nil, b == nil, len(a), len(b))
	ja, _ := json.Marshal(a)
	jb, _ := json.Marshal(b)
	fmt.Println(string(ja), string(jb))
	// Output:
	// true false 0 0
	// null []
}
