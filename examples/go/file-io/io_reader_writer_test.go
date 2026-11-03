package fileio

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type upperWriter struct{ w io.Writer }

func (u upperWriter) Write(p []byte) (int, error) { // any type with Write is an io.Writer
	return u.w.Write(bytes.ToUpper(p))
}

func Example_readerWriter() {
	var sb strings.Builder
	io.Copy(upperWriter{&sb}, strings.NewReader("go io"))
	fmt.Println(sb.String())

	r := io.LimitReader(strings.NewReader("abcdef"), 3) // readers compose
	b, _ := io.ReadAll(r)
	fmt.Println(string(b))
	// Output:
	// GO IO
	// abc
}
