//go:build cgo

package unsafecgo

import "testing"

func BenchmarkGoAdd(b *testing.B) {
	add := func(a, b int) int { return a + b }
	for i := 0; i < b.N; i++ {
		_ = add(i, 1)
	}
}

func BenchmarkCgoAdd(b *testing.B) { // each call crosses the Go/C boundary
	for i := 0; i < b.N; i++ {
		_ = Add(i, 1)
	}
}
