//go:build cgo

package unsafecgo

/*
#include <stdlib.h>
#include <string.h>

static int add(int a, int b) { return a + b; }

static size_t c_strlen(const char *s) { return strlen(s); }
*/
import "C"

import "unsafe"

// Add calls the C function add.
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

// CStrlen passes a Go string to C: copy it to C memory, then free it.
func CStrlen(s string) int {
	cs := C.CString(s) // malloc'd by C, not tracked by the Go GC
	defer C.free(unsafe.Pointer(cs))
	return int(C.c_strlen(cs))
}
