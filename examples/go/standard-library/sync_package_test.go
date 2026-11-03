package standardlibrary

import (
	"fmt"
	"sync"
)

func Example_syncMap() {
	var m sync.Map // concurrent map for write-once / read-many keys
	m.Store("a", 1)
	v, ok := m.Load("a")
	actual, loaded := m.LoadOrStore("a", 2) // keeps the existing value
	fmt.Println(v, ok, actual, loaded)
	// Output: 1 true 1 true
}
