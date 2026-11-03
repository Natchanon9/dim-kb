package unsafecgo

import (
	"fmt"
	"unsafe"
)

func Example_unsafePackage() {
	// sizes on a 64-bit platform: string = ptr+len, slice = ptr+len+cap
	fmt.Println(unsafe.Sizeof(int64(0)), unsafe.Sizeof(""), unsafe.Sizeof([]int{}))
	// Output: 8 16 24
}
