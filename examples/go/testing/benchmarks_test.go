package gotesting

import (
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	parts := []string{"a", "b", "c"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ { // Go 1.24+: for b.Loop() { ... }
		strings.Join(parts, ",")
	}
}
