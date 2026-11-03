package structs

import (
	"errors"
	"fmt"
	"time"
)

type Server struct {
	addr    string
	timeout time.Duration
}

func NewServer(addr string) (*Server, error) {
	if addr == "" {
		return nil, errors.New("addr required")
	}
	return &Server{addr: addr, timeout: 30 * time.Second}, nil // set defaults here
}

func Example_constructor() {
	s, _ := NewServer(":8080")
	_, err := NewServer("")
	fmt.Println(s.addr, s.timeout, err)
	// Output: :8080 30s addr required
}
