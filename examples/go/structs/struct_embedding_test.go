package structs

import "fmt"

type Base struct{ ID int }

func (b Base) Describe() string { return fmt.Sprintf("id=%d", b.ID) }

type Product struct {
	Base // embedded: fields and methods are promoted
	Name string
}

func Example_embedding() {
	p := Product{Base{7}, "pen"}
	fmt.Println(p.ID, p.Describe(), p.Base.ID)
	// Output: 7 id=7 7
}
