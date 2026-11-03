package fileio

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Example_os() {
	dir, _ := os.MkdirTemp("", "demo")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "note.txt")
	if err := os.WriteFile(path, []byte("hello\n"), 0o644); err != nil {
		panic(err)
	}
	data, _ := os.ReadFile(path)
	fmt.Print(string(data))

	_, err := os.Stat(filepath.Join(dir, "missing"))
	fmt.Println(errors.Is(err, os.ErrNotExist))
	// Output:
	// hello
	// true
}
