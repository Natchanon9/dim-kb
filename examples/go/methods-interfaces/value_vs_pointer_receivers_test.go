package methodsinterfaces

import "fmt"

type Counter struct{ n int }

func (c Counter) Value() int { return c.n }
func (c *Counter) Inc()      { c.n++ } // changes the caller's Counter
func (c Counter) IncOnCopy() { c.n++ } // changes a copy: no effect

func Example_receivers() {
	var c Counter
	c.Inc() // Go takes &c for you because c is addressable
	c.Inc()
	c.IncOnCopy()
	fmt.Println(c.Value())
	// Output: 2
}
