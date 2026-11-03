package methodsinterfaces

import (
	"bytes"
	"fmt"
	"io"
)

type Logger interface{ Log(msg string) }

type ConsoleLogger struct{ prefix string }

func (c ConsoleLogger) Log(msg string) { fmt.Println(c.prefix + msg) }

var _ Logger = ConsoleLogger{} // compile-time check, optional

func start(l Logger) { l.Log("started") }

func Example_implicit() {
	start(ConsoleLogger{"[app] "})

	var w io.Writer = &bytes.Buffer{} // bytes.Buffer never mentions io.Writer
	fmt.Fprint(w, "x")
	fmt.Println(w.(*bytes.Buffer).Len())
	// Output:
	// [app] started
	// 1
}
