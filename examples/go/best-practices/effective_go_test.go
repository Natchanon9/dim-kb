package bestpractices

import "fmt"

// Buffer's zero value is ready to use: no constructor needed.
type Buffer struct{ data []byte }

func (b *Buffer) Write(p []byte) { b.data = append(b.data, p...) }

func (b *Buffer) String() string { return string(b.data) }

func Example_effectiveGo() {
	var b Buffer
	b.Write([]byte("go"))
	fmt.Println(b.String())
	// Output: go
}
