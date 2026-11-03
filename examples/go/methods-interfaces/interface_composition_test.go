package methodsinterfaces

import "fmt"

type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface { // composed from smaller interfaces
	Reader
	Writer
}

type memFile struct{ data []string }

func (m *memFile) Write(s string) { m.data = append(m.data, s) }
func (m *memFile) Read() string {
	if len(m.data) == 0 {
		return ""
	}
	s := m.data[0]
	m.data = m.data[1:]
	return s
}

func Example_composition() {
	var rw ReadWriter = &memFile{}
	rw.Write("hi")
	fmt.Println(rw.Read())
	// Output: hi
}
