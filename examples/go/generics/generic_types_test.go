package generics

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack[T]) Len() int { return len(s.items) }

func Example_types() {
	var s Stack[string] // instantiate with a concrete type
	s.Push("a")
	s.Push("b")
	v, ok := s.Pop()
	fmt.Println(v, ok, s.Len())
	// Output: b true 1
}
