package concurrencypatterns

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) }, // used when the pool is empty
}

func render(name string) string {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset() // a pooled object may hold old data
	buf.WriteString("hello, ")
	buf.WriteString(name)
	return buf.String()
}

func Example_pool() {
	fmt.Println(render("go"))
	fmt.Println(render("pool"))
	// Output:
	// hello, go
	// hello, pool
}
