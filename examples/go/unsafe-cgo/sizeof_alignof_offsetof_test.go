package unsafecgo

import (
	"fmt"
	"unsafe"
)

type padded struct {
	A bool  // 1 byte + 7 padding
	B int64 // 8
	C bool  // 1 byte + 7 padding
}

type packed struct {
	B int64 // 8
	A bool  // 1
	C bool  // 1 + 6 padding
}

func Example_layout() {
	fmt.Println(unsafe.Sizeof(padded{}), unsafe.Sizeof(packed{}))
	fmt.Println(unsafe.Alignof(int64(0)), unsafe.Offsetof(padded{}.B))
	// Output:
	// 24 16
	// 8 8
}
