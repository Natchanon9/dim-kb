// Package slices shows the append/capacity behaviour summarised in
// src/content/docs/programming/go/arrays-slices/append-copy.md.
package slices

// Grow appends n values and returns the slice plus every capacity it went through.
func Grow(n int) ([]int, []int) {
	var s []int
	caps := []int{cap(s)}
	for i := range n {
		s = append(s, i)
		if c := cap(s); c != caps[len(caps)-1] {
			caps = append(caps, c)
		}
	}
	return s, caps
}

// Clone copies src into a new backing array so changes don't leak back.
func Clone(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
