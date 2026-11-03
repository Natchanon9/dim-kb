package bestpractices

import (
	"fmt"
	"time"
)

type Clock interface{ Now() time.Time }

type Greeter struct{ clock Clock } // dependencies are fields

func NewGreeter(c Clock) *Greeter { return &Greeter{clock: c} } // injected through the constructor

func (g *Greeter) Greet() string {
	if g.clock.Now().Hour() < 12 {
		return "good morning"
	}
	return "good afternoon"
}

type fixedClock time.Time

func (f fixedClock) Now() time.Time { return time.Time(f) }

func Example_dependencyInjection() {
	morning := fixedClock(time.Date(2026, 1, 1, 9, 0, 0, 0, time.UTC))
	fmt.Println(NewGreeter(morning).Greet())
	// Output: good morning
}
