package slices

import "testing"

func TestGrow(t *testing.T) {
	s, caps := Grow(5)
	if len(s) != 5 {
		t.Fatalf("len = %d, want 5", len(s))
	}
	if caps[0] != 0 {
		t.Errorf("nil slice cap = %d, want 0", caps[0])
	}
}

func TestCloneIsIndependent(t *testing.T) {
	a := []int{1, 2, 3}
	b := Clone(a)
	b[0] = 99
	if a[0] != 1 {
		t.Errorf("original changed to %d; Clone must not share the backing array", a[0])
	}
}
