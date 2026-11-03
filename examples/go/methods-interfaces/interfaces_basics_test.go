package methodsinterfaces

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "woof" }

type Robot struct{ ID int }

func (r Robot) Speak() string { return fmt.Sprintf("beep %d", r.ID) }

func Example_interfaces() {
	for _, s := range []Speaker{Dog{}, Robot{7}} {
		fmt.Println(s.Speak())
	}
	// Output:
	// woof
	// beep 7
}
