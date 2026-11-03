package pointers

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Sum() int {
	if n == nil { // a nil receiver is legal; guard it
		return 0
	}
	return n.Val + n.Next.Sum()
}

func Example_nilPointers() {
	list := &Node{1, &Node{2, nil}}
	var empty *Node
	fmt.Println(list.Sum(), empty.Sum())
	// Output: 3 0
}
