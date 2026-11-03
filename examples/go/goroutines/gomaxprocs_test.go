package goroutines

import (
	"runtime"
	"testing"
)

func TestGOMAXPROCS(t *testing.T) {
	prev := runtime.GOMAXPROCS(2) // set; returns the previous value
	defer runtime.GOMAXPROCS(prev)
	if got := runtime.GOMAXPROCS(0); got != 2 { // 0 = query only
		t.Fatalf("GOMAXPROCS = %d, want 2", got)
	}
}
